package errf

import (
	"fmt"
)

// DefaultErrflow is default Errflow instance.
// Clients should never set its value.
var DefaultErrflow = &Errflow{}

// Errflow contains configuration for error hanlding logic.
// It exposes an immutable API for clients, but it is not thread-safe.
type Errflow struct {
	wrapper func(err error) error
	logStrategy
	returnStrategy

	deferredOptions []ErrflowOption
	appliedOptions  []ErrflowOption
}

func (ef *Errflow) applyDeferredOptions() {
	newEf := ef
	for _, option := range ef.deferredOptions {
		newEf = option(newEf)
	}
	*ef = *newEf
	ef.appliedOptions = append(ef.appliedOptions, ef.deferredOptions...)
	ef.deferredOptions = nil
}

func (ef *Errflow) copy() *Errflow {
	return &Errflow{
		wrapper:        ef.wrapper,
		logStrategy:    ef.logStrategy,
		returnStrategy: ef.returnStrategy,

		deferredOptions: ef.deferredOptions,
		appliedOptions:  ef.appliedOptions,
	}
}

// ErrflowOption is used for extending config API.
// Clients can extend config API with new methods by implementing ErrflowOption.
//
// Example:
//  func WrapInOurError(ef *Errflow) *Errflow {
//  	return ef.With(errf.Wrapper(func (err error) error {
//  		return ourerrors.Wrap(err)
//  	}))
//  }
//
//  func exampleUsage() (err error) {
//  	defer errf.IfError().Apply(WrapInOurError).ThenAssignTo(&err)
//  	// ...
//  }
type ErrflowOption func(ef *Errflow) *Errflow

// With adds additional configs to Errflow instance.
// It returns a new instance. Original instance is unmodified.
func (ef *Errflow) With(options ...ErrflowOption) *Errflow {
	if ef == nil {
		ef = DefaultErrflow
	}
	result := ef.copy()
	result.deferredOptions = append([]ErrflowOption{}, ef.deferredOptions...)
	result.deferredOptions = append(result.deferredOptions, options...)
	return result
}

// Opts returns all options, which were applied on top of DefaultErrflow instance.
func (ef *Errflow) Opts() []ErrflowOption {
	if ef.appliedOptions == nil {
		return ef.deferredOptions
	}

	var result []ErrflowOption
	result = append(result, ef.appliedOptions...)
	result = append(result, ef.deferredOptions...)
	return result
}

// With is an alias for DefaultErrflow.With(...).
func With(options ...ErrflowOption) *Errflow {
	return DefaultErrflow.With(options...)
}

type errflowThrowItem struct {
	ef  *Errflow
	err error
}

type errflowThrow struct {
	items []errflowThrowItem
}

// ImplementTry is used to implement a strongly-typed errflow.Try(...) for new types.
//
// Example:
//  package fancypackage
//
//  var Errf = FancyPackageErrflow{}
//
//  type FancyPackageErrflow struct {
//  	errflow *Errflow
//  }
//
//  func (ef FancyPackageErrflow) With(options ...ErrflowOption) FancyPackageErrflow {
//  	return FancyPackageErrflow{errflow: errf.errflow.With(options...)}
//  }
//
//  func (ef FancyPackageErrflow) TryCustomType1(value *CustomType1, err error) *CustomType1 {
//  	errf.errflow.ImplementTry(recover(), err)
//  	return value
//  }
//
//  func (ef FancyPackageErrflow) TryCustomType2(value *CustomType2, err error) *CustomType2 {
//  	errf.errflow.ImplementTry(recover(), err)
//  	return value
//  }
//
//  package main
//
//  func ProcessCustomStruct() (err error) {
//  	defer errflow.IfError().ThenAssignTo(&err)
//
//  	customStructValue := fancypackage.Errf.TryCustomStruct(
//  		fancypackage.ReadCustomStruct())
//
//  	// ...
//  }
func (ef *Errflow) ImplementTry(recoverObj interface{}, err error) error {
	if ef == nil {
		ef = DefaultErrflow
	}
	globalErrflowValidator.validate()

	var errflowThrowObj errflowThrow
	if recoverObj != nil {
		recoveredErrflowThrowObj, ok := recoverObj.(errflowThrow)
		if ok {
			errflowThrowObj.items = append(errflowThrowObj.items, recoveredErrflowThrowObj.items...)
		} else {
			panic(recoverObj)
		}
	}
	if err != nil {
		errflowThrowObj.items = append(errflowThrowObj.items, errflowThrowItem{
			ef:  ef,
			err: err,
		})
	}
	if len(errflowThrowObj.items) > 0 {
		panic(errflowThrowObj)
	}
	return err
}

// TryErr sends error to IfError() handler for processing, if there is an error.
//
// It is required that 'defer errf.IfError().Then...' is configured in the same
// function as TryErr, otherwise validation will fail when running tests.
//
// TryErr always returns nil, but type system allows using it to skip return nil statement:
//   errflow.TryErr(functionCall())
//   return nil
// is the same as:
//   return errflow.TryErr(functionCall())
func (ef *Errflow) TryErr(err error) error {
	return ef.ImplementTry(recover(), err)
}

// TryErr is an alias for DefaultErrflow.TryErr(...).
func TryErr(err error) error {
	return DefaultErrflow.ImplementTry(recover(), err)
}

// TryAny sends error to Catcher for processing, if there is an error.
// If there is no error, it returns value as a generic interface{}.
//
// Example:
//  function ProcessFile() (err error) {
//    defer errf.IfError().ThenAssignTo(&err)
//
//    file := errf.TryAny(os.Create("file.go")).(*os.File)
//    defer errf.TryErr(file.Close())
//
//    // Write to file ...
//  }
//
// Tip: prefer using typed functions, defined in either this library, or
// custom ones, implemented using errf.ImplementTry(...).
//
// Example above can usually rewritten as:
//  function ProcessFile() (err error) {
//    defer errf.IfError().ThenAssignTo(&err)
//
//    writer := errf.Io.TryWriteCloser(os.Create("file.go"))
//    defer errf.TryErr(writer.Close())
//
//    // Write to file ...
//  }
func (ef *Errflow) TryAny(value interface{}, err error) interface{} {
	ef.ImplementTry(recover(), err)
	return value
}

// TryAny is an alias for DefaultErrflow.TryAny(...).
func TryAny(value interface{}, err error) interface{} {
	DefaultErrflow.ImplementTry(recover(), err)
	return value
}

// TryDiscard sends error to IfError() handler for processing, if there is an error.
// Non-error value returned from a function is discarded.
//
// Example:
//  function writeBuf(w io.Writer, buf []byte) (err error) {
//    defer errf.IfError().ThenAssignTo(&err)
//
//    return errf.TryDiscard(w.Write(buf))
//  }
func (ef *Errflow) TryDiscard(value interface{}, err error) error {
	return ef.ImplementTry(recover(), err)
}

// TryDiscard is an alias for DefaultErrflow.TryDiscard(...).
func TryDiscard(value interface{}, err error) error {
	return DefaultErrflow.ImplementTry(recover(), err)
}

// TryCondition creates and sends error to IfError() handler for processing, if condition is true.
func (ef *Errflow) TryCondition(condition bool, format string, a ...interface{}) error {
	if condition {
		return ef.ImplementTry(recover(), fmt.Errorf(format, a...))
	}
	return nil
}

// TryCondition is an alias for DefaultErrflow.TryCondition(...).
func TryCondition(condition bool, format string, a ...interface{}) error {
	if condition {
		return DefaultErrflow.ImplementTry(recover(), fmt.Errorf(format, a...))
	}
	return nil
}

// Log is an alias for errf.Log(...) function.
func (ef *Errflow) Log(err error) error {
	return Log(err)
}

// Log logs error, if not nil.
// Doesn't affect control flow.
func Log(err error) error {
	if err != nil {
		globalLogFn(&LogMessage{
			Format: "%s",
			A:      []interface{}{err.Error()},
			Stack:  getStringErrorStackTraceFn(),
			Tags:   []string{"errflow", "error"},
		})
	}
	return err
}
