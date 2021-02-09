package errflow

import "fmt"

type errflowThrow struct {
	errs []error
}

// ImplementCheck is used to implement a strongly-typed errflow.Check(...)
// for processing function return values for custom types.
//
// Example:
//   package fancypackage
//
//   type CustomStruct struct { ... }
//
//   func ErrflowCustomStruct(value *CustomStruct, err error) *CustomStruct {
//     ImplementCheck(recover(), err)
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
func ImplementCheck(recoverObj interface{}, err error) error {
	globalErrflowValidator.validate()

	var errs []error
	if recoverObj != nil {
		errflowThrowObj, ok := recoverObj.(errflowThrow)
		if ok {
			errs = errflowThrowObj.errs
		} else {
			panic(recoverObj)
		}
	}
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		panic(errflowThrow{errs: errs})
	}
	return err
}

// Check sends error to Catcher for processing, if there is an error.
//
// It is required that 'defer errflow.Catch()' is configured in the same
// function as Check, otherwise validation will fail when running tests.
//
// Check always returns nil, but type system allows using it to skip
// return nil statement:
//   errflow.Check(functionCall())
//   return nil
// is the same as:
//   return errflow.Check(functionCall())
func Check(err error) error {
	return ImplementCheck(recover(), err)
}

// CheckAny sends error to Catcher for processing, if there is an error.
// If there is no error, it returns value as a generic interface{}.
//
// Example:
//  function ProcessFile() (err error) {
//    defer errflow.IfError().ThenAssignTo(&err)
//
//    file := errflow.CheckAny(os.Create("file.go")).(*os.File)
//    defer errflow.Check(file.Close())
//
//    // Write to file ...
//  }
//
// Tip: prefer using typed functions, defined in either this library, or
// custom ones, implemented using errflow.ImplementCheck(...).
//
// Example above can usually rewritten as:
//  function ProcessFile() (err error) {
//    defer errflow.IfError().ThenAssignTo(&err)
//
//    writer := errflow.CheckIoWriteCloser(os.Create("file.go"))
//    defer errflow.Check(writer.Close())
//
//    // Write to file ...
//  }
func CheckAny(value interface{}, err error) interface{} {
	ImplementCheck(recover(), err)
	return value
}

// CheckDiscard sends error to Catcher for processing, if there is an error.
// Non-error value returned from a function is ignored.
//
// Example:
//  function writeBuf(w io.Writer, buf []byte) (err error) {
//    defer errflow.IfError().ThenAssignTo(&err)
//
//    return errflow.CheckDiscard(w.Write(buf))
//  }
func CheckDiscard(value interface{}, err error) error {
	return ImplementCheck(recover(), err)
}

// OnlyLog logs error, if not nil, but doesn't affect control flow.
func OnlyLog(err error) error {
	if err != nil {
		globalLogFn(err.Error())
	}
	return err
}

// ErrorIf sends error to Catcher for processing, if condition is true.
func ErrorIf(condition bool, format string, a ...interface{}) {
	if condition {
		ImplementCheck(recover(), fmt.Errorf(format, a...))
	}
}
