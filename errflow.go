package errf

import (
	"fmt"
)

var DefaultErrflow = &Errflow{}
var Io = IoErrflow{}
var Bufio = BufioErrflow{}
var Std = StdErrflow{}

type Errflow struct {
	wrapper func(err error) error
	logStrategy
	returnStrategy

	deferredOptions []ErrflowOption
}

func (ef *Errflow) applyDeferredOptions() {
	newEf := ef
	for _, option := range ef.deferredOptions {
		newEf = option(newEf)
	}
	*ef = *newEf
	ef.deferredOptions = nil
}

func (ef *Errflow) copy() *Errflow {
	return &Errflow{
		wrapper:        ef.wrapper,
		logStrategy:    ef.logStrategy,
		returnStrategy: ef.returnStrategy,

		deferredOptions: ef.deferredOptions,
	}
}

type ErrflowOption func(errflowOptions *Errflow) *Errflow

func (ef *Errflow) With(options ...ErrflowOption) *Errflow {
	if ef == nil {
		ef = DefaultErrflow
	}
	result := ef.copy()
	result.deferredOptions = append([]ErrflowOption{}, ef.deferredOptions...)
	result.deferredOptions = append(result.deferredOptions, options...)
	return result
}

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

// ImplementTry is used to implement a strongly-typed errflow.Try(...)
// for processing function return values for custom types.
//
// Example:
//   package fancypackage
//
//   type CustomStruct struct { ... }
//
//   func ErrflowCustomStruct(value *CustomStruct, err error) *CustomStruct {
//     ImplementTry(recover(), err)
//     return value
//   }
//
//   func ReadCustomStruct() (*CustomStruct, error) { ... }
//
//
//   package main
//
//   func ProcessCustomStruct() (err error) {
//     defer errflow.IfError().ThenAssignTo(&err)
//
//     customStruct := fancypackage.ErrflowCustomStruct(fancypackage.ReadCustomStruct())
//
//     // ...
//   }
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

// Try sends error to Catcher for processing, if there is an error.
//
// It is required that 'defer errflow.Catch()' is configured in the same
// function as Try, otherwise validation will fail when running tests.
//
// Try always returns nil, but type system allows using it to skip
// return nil statement:
//   errflow.Try(functionCall())
//   return nil
// is the same as:
//   return errflow.Try(functionCall())
func (ef *Errflow) TryErr(err error) error {
	return ef.ImplementTry(recover(), err)
}

func TryErr(err error) error {
	return DefaultErrflow.ImplementTry(recover(), err)
}

// TryAny sends error to Catcher for processing, if there is an error.
// If there is no error, it returns value as a generic interface{}.
//
// Example:
//  function ProcessFile() (err error) {
//    defer errflow.IfError().ThenAssignTo(&err)
//
//    file := errflow.TryAny(os.Create("file.go")).(*os.File)
//    defer errflow.Try(file.Close())
//
//    // Write to file ...
//  }
//
// Tip: prefer using typed functions, defined in either this library, or
// custom ones, implemented using errflow.ImplementTry(...).
//
// Example above can usually rewritten as:
//  function ProcessFile() (err error) {
//    defer errflow.IfError().ThenAssignTo(&err)
//
//    writer := errflow.TryIoWriteCloser(os.Create("file.go"))
//    defer errflow.Try(writer.Close())
//
//    // Write to file ...
//  }
func (ef *Errflow) TryAny(value interface{}, err error) interface{} {
	ef.ImplementTry(recover(), err)
	return value
}

func TryAny(value interface{}, err error) interface{} {
	DefaultErrflow.ImplementTry(recover(), err)
	return value
}

// TryDiscard sends error to Catcher for processing, if there is an error.
// Non-error value returned from a function is ignored.
//
// Example:
//  function writeBuf(w io.Writer, buf []byte) (err error) {
//    defer errflow.IfError().ThenAssignTo(&err)
//
//    return errflow.TryDiscard(w.Write(buf))
//  }
func (ef *Errflow) TryDiscard(value interface{}, err error) error {
	return ef.ImplementTry(recover(), err)
}

func TryDiscard(value interface{}, err error) error {
	return DefaultErrflow.ImplementTry(recover(), err)
}

func (ef *Errflow) TryCondition(condition bool, format string, a ...interface{}) error {
	if condition {
		return ef.ImplementTry(recover(), fmt.Errorf(format, a...))
	}
	return nil
}

// TryCondition creates and sends error to Catcher for processing, if condition is true.
func TryCondition(condition bool, format string, a ...interface{}) error {
	if condition {
		return DefaultErrflow.ImplementTry(recover(), fmt.Errorf(format, a...))
	}
	return nil
}

// Log logs error, if not nil. Doesn't affect control flow.
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
