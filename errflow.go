package errflow

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
//     defer errflow.Catch().WriteTo(&err)
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
	return nil
}

// Check sends error to Catcher for processing, if there is an error.
// It is required that 'defer errflow.Catch()' is configured in the same
// function as Check, otherwise validation will fail when running tests.
//
// Tip: prefer a shortcut function C() instead of a full Check().
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

// C is a shortcut for Catch.
// Check sends error to Catcher for processing, if there is an error.
// It is required that 'defer errflow.Catch()' is configured in the same
// function as Check, otherwise validation will fail when running tests.
//
// Tip: prefer a shortcut function C() instead of a full Check().
//
// Check always returns nil, but type system allows using it to skip
// return nil statement:
//   errflow.Check(functionCall())
//   return nil
// is the same as:
//   return errflow.Check(functionCall())
func C(err error) error {
	return ImplementCheck(recover(), err)
}

// Untyped sends error to Catcher for processing, if there is an error.
// If there is no error, it returns value as a generic interface.
//
// Tip: prefer a shortcut function U() instead of a full Untyped().
//
// Example:
//  function ProcessFile() (err error) {
//    defer errflow.Catch().WriteTo(&err)
//
//    file := errflow.Untyped(os.Create("file.go")).(*os.File)
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
//    defer errflow.Catch().WriteTo(&err)
//
//    writer := errflow.IoWriteCloser(os.Create("file.go"))
//    defer errflow.Check(writer.Close())
//
//    // Write to file ...
//  }
func Untyped(value interface{}, err error) interface{} {
	ImplementCheck(recover(), err)
	return value
}

// U is a shortcut for Untyped.
// Untyped sends error to Catcher for processing, if there is an error.
// If there is no error, it returns value as a generic interface.
//
// Tip: prefer a shortcut function U() instead of a full Untyped().
//
// Example:
//  function ProcessFile() (err error) {
//    defer errflow.Catch().WriteTo(&err)
//
//    file := errflow.U(os.Create("file.go")).(*os.File)
//    defer errflow.C(file.Close())
//
//    // Write to file ...
//  }
//
// Tip: prefer using typed functions, defined in either this library, or
// custom ones, implemented using errflow.ImplementCheck(...).
//
// Example above can usually rewritten as:
//  function ProcessFile() (err error) {
//    defer errflow.Catch().WriteTo(&err)
//
//    writer := errflow.IoWriteCloser(os.Create("file.go"))
//    defer errflow.C(writer.Close())
//
//    // Write to file ...
//  }
func U(value interface{}, err error) interface{} {
	ImplementCheck(recover(), err)
	return value
}

// IgnoreReturnValue sends error to Catcher for processing, if there is an error.
// Non-error value returned from a function is ignored.
//
// Example:
//  function writeBuf(w io.Writer, buf []byte) (err error) {
//    defer errflow.Catch().WriteTo(&err)
//
//    errflow.IgnoreReturnValue(w.Write(buf))
//  }
//
// Tip: prefer a shortcut function I() instead of a full IgnoreReturnValue().
func IgnoreReturnValue(value interface{}, err error) error {
	return ImplementCheck(recover(), err)
}

// I is a shortcut for IgnoreReturnValue.
// IgnoreReturnValue sends error to Catcher for processing, if there is an error.
// Non-error value returned from a function is ignored.
//
// Example:
//  function writeBuf(w io.Writer, buf []byte) (err error) {
//    defer errflow.Catch().WriteTo(&err)
//
//    errflow.I(w.Write(buf))
//  }
//
// Tip: prefer a shortcut function I() instead of a full IgnoreReturnValue().
func I(value interface{}, err error) error {
	return ImplementCheck(recover(), err)
}
