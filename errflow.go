package errf

import (
	"fmt"
)

// DefaultErrflow is default Errflow instance.
// Clients should never set its value.
var DefaultErrflow = &Errflow{}

// Errflow contains configuration for error handling logic.
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
	errflow := ef
	if errflow == nil {
		errflow = DefaultErrflow
	}
	result := errflow.copy()
	result.deferredOptions = append([]ErrflowOption{}, errflow.deferredOptions...)
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

// AsOpts returns ErrflowOption copies all configs of Errflow.
func (ef *Errflow) AsOpts() ErrflowOption {
	return OptsFrom(ef)
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

// CheckResult default object is returned by Check* functions
// which don't return any value, as a helper for return statements:
//  return errf.CheckErr(err).IfOkReturnNil
type CheckResult struct {
	IfOkReturnNil error
}

// ImplementCheck is used to implement a strongly-typed errflow.Check(...) for new types.
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
//  	return FancyPackageErrflow{errflow: ef.errflow.With(options...)}
//  }
//
//  func (ef FancyPackageErrflow) CheckCustomType1(value *CustomType1, err error) *CustomType1 {
//  	ef.errflow.ImplementCheck(recover(), err)
//  	return value
//  }
//
//  func (ef FancyPackageErrflow) CheckCustomType2(value *CustomType2, err error) *CustomType2 {
//  	ef.errflow.ImplementCheck(recover(), err)
//  	return value
//  }
//
//  package main
//
//  func ProcessCustomStruct() (err error) {
//  	defer errflow.IfError().ThenAssignTo(&err)
//
//  	customStructValue := fancypackage.Errf.CheckCustomType1(
//  		fancypackage.ReadCustomType1())
//
//  	// ...
//  }
func (ef *Errflow) ImplementCheck(recoverObj interface{}, err error) CheckResult {
	errflow := ef
	if errflow == nil {
		errflow = DefaultErrflow
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
			ef:  errflow,
			err: err,
		})
	}
	if len(errflowThrowObj.items) > 0 {
		panic(errflowThrowObj)
	}
	return CheckResult{}
}

// CheckErr sends error to IfError() handler for processing, if there is an error.
//
// It is required that 'defer errf.IfError().Then...' is configured in the same
// function as CheckErr, otherwise validation will fail when running tests.
//
// CheckErr always returns nil, but type system allows using it to skip return nil statement:
//   errflow.CheckErr(functionCall())
//   return nil
// is the same as:
//   return errflow.CheckErr(functionCall())
func (ef *Errflow) CheckErr(err error) CheckResult {
	return ef.ImplementCheck(recover(), err)
}

// CheckErr is an alias for DefaultErrflow.CheckErr(...).
func CheckErr(err error) CheckResult {
	return DefaultErrflow.ImplementCheck(recover(), err)
}

// CheckDeferErr calls closeFn and checks for return error.
// Useful in defer statements:
//  writer := ...
//  defer errf.CheckDeferErr(writer.Close)
func (ef *Errflow) CheckDeferErr(closeFn func() error) CheckResult {
	return ef.ImplementCheck(recover(), closeFn())
}

// CheckDeferErr is an alias for DefaultErrflow.CheckDeferErr(...).
func CheckDeferErr(closeFn func() error) CheckResult {
	return DefaultErrflow.ImplementCheck(recover(), closeFn())
}

// CheckAny sends error to IfError() handler for processing, if there is an error.
// If there is no error, it returns value as a generic interface{}.
//
// Example:
//  function ProcessFile() (err error) {
//    defer errf.IfError().ThenAssignTo(&err)
//
//    file := errf.CheckAny(os.Create("file.go")).(*os.File)
//    defer errf.CheckDeferErr(file.Close)
//
//    // Write to file ...
//  }
//
// Tip: prefer using typed functions, defined in either this library, or
// custom ones, implemented using errf.ImplementCheck(...).
//
// Example above can usually rewritten as:
//  function ProcessFile() (err error) {
//    defer errf.IfError().ThenAssignTo(&err)
//
//    writer := errf.Io.CheckWriteCloser(os.Create("file.go"))
//    defer errf.CheckDeferErr(writer.Close)
//
//    // Write to file ...
//  }
func (ef *Errflow) CheckAny(value interface{}, err error) interface{} {
	ef.ImplementCheck(recover(), err)
	return value
}

// CheckAny is an alias for DefaultErrflow.CheckAny(...).
func CheckAny(value interface{}, err error) interface{} {
	DefaultErrflow.ImplementCheck(recover(), err)
	return value
}

// CheckDiscard sends error to IfError() handler for processing, if there is an error.
// Non-error value returned from a function is discarded.
//
// Example:
//  function writeBuf(w io.Writer, buf []byte) (err error) {
//    defer errf.IfError().ThenAssignTo(&err)
//
//    return errf.CheckDiscard(w.Write(buf))
//  }
func (ef *Errflow) CheckDiscard(_ interface{}, err error) CheckResult {
	return ef.ImplementCheck(recover(), err)
}

// CheckDiscard is an alias for DefaultErrflow.CheckDiscard(...).
func CheckDiscard(_ interface{}, err error) CheckResult {
	return DefaultErrflow.ImplementCheck(recover(), err)
}

// CheckCondition creates and sends error to IfError() handler for processing, if condition is true.
func (ef *Errflow) CheckCondition(condition bool, format string, a ...interface{}) CheckResult {
	if condition {
		return ef.ImplementCheck(recover(), fmt.Errorf(format, a...))
	}
	return CheckResult{}
}

// CheckCondition is an alias for DefaultErrflow.CheckCondition(...).
func CheckCondition(condition bool, format string, a ...interface{}) CheckResult {
	if condition {
		return DefaultErrflow.ImplementCheck(recover(), fmt.Errorf(format, a...))
	}
	return CheckResult{}
}

// CheckAssert creates and sends error to IfError() handler for processing, if condition is false.
func (ef *Errflow) CheckAssert(condition bool, format string, a ...interface{}) CheckResult {
	if !condition {
		return ef.ImplementCheck(recover(), fmt.Errorf(format, a...))
	}
	return CheckResult{}
}

// CheckAssert is an alias for DefaultErrflow.CheckAssert(...).
func CheckAssert(condition bool, format string, a ...interface{}) CheckResult {
	if !condition {
		return DefaultErrflow.ImplementCheck(recover(), fmt.Errorf(format, a...))
	}
	return CheckResult{}
}

// Log logs error, if not nil.
// Always logs, even if log strategy is LogStrategyNever.
// Doesn't affect control flow.
func (ef *Errflow) Log(err error) {
	if err != nil {
		ef.applyDeferredOptions()
		if ef.wrapper != nil {
			err = ef.wrapper(err)
		}
		if err == nil {
			panic("error wrapper returned nil error")
		}
		globalLogFn(&LogMessage{
			Format: "%s",
			A:      []interface{}{err.Error()},
			Stack:  getStringErrorStackTraceFn(),
			Tags:   []string{"errorflow", "error"},
		})
	}
}

// Log is an alias for DefaultErrflow.Log(...).
func Log(err error) {
	DefaultErrflow.Log(err)
}

// LogDefer calls closeFn, then calls Log(...) on result of a call.
// Useful in defer statements:
//  reader := ...
//  defer errf.LogDefer(reader.Close)
func (ef *Errflow) LogDefer(closeFn func() error) {
	ef.Log(closeFn())
}

// LogDefer is an alias for DefaultErrflow.LogDefer(...).
func LogDefer(closeFn func() error) {
	DefaultErrflow.LogDefer(closeFn)
}
