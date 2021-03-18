// Package errf provides declarative error handling for Go.
//
// Basic usage
//
// Each use of errorflow has a scope of a single function.
//
// If function uses errorflow its first statement should be IfError() handler
// configuration:
//
//  func exampleFunction() (err error) {
//  	defer errf.IfError().ReturnFirst().LogIfSuppressed().ThenAssignTo(&err)
//
//  	// ... business logic ...
//  }
//
// Default configuration is to return first error and never log errors.
//
//  defer errf.IfError().ThenAssignTo(&err)
//  	/* same as */
//  defer errf.IfError().ReturnFirst().LogNever().ThenAssignTo(&err)
//
// IfError() handler should always be terminated by one of Then*() functions.
//
// Then, when function has IfError() handler setup, all error checks should be
// done via one of Check* functions:
//
//  func exampleFunction() (err error) {
//  	defer errf.IfError().ThenAssignTo(&err)
//
//  	// Just checking err.
//  	err := someFunction1()
//  	errf.CheckErr(err)
//  		/* same as */
//  	errf.CheckErr(someFunction1())
//
//  	// Checking err, discarding return value.
//  	_, err := someFunction2()
//  	errf.CheckErr(err)
//  		/* same as */
//  	errf.CheckDiscard(someFunction2())
//
//  	// Checking err using return value (untyped).
//  	var value interface{}
//  	value, err = someFunction3()
//  	errf.CheckErr(err)
//  	typedValue := value.(type)
//  		/* same as */
//  	typedValue := errf.CheckAny(someFunction3()).(type)
//  	// NOTE: using CheckAny is not recommended,
//  	// clients should prefer using either standard typed function
//  	// or create custom ones (see 'Extending errorflow for custom types' below).
//
//  	// Checking err using return value (typed).
//  	int64Value, err := someFunction4()
//  	errf.CheckErr(err)
//  		/* same as */
//  	int64Value := errf.Std.CheckInt64(someFunction4())
//  		/* same as */
//  	int64Value, err := errf.Std.CheckInt64Err(someFunction4())
//
//  	// It is not recommended to manually return errors (return fmt.Errorf(...))
//  	// or assign errors (err = fmt.Errorf(...)) for functions with IfError() handler.
//  	// In such cases, wrap error in errf.CheckErr(...) to make sure it is handled
//  	// by IfError() handler.
//  	return errf.CheckErr(someFunction5())
//  }
//
// When any of Check* functions encounters an error, it immediately sends a program flow control
// to IfError() handler, unwinding stack of already registered 'defers'. Internally it uses
// panic to accomplish this goal.
//
// Never use Check* functions in functions without IfError() handler set up
// (including nested anonymous functions).
//
// Usage correctness validation
//
// When running Go tests, errorflow automatically verifies correctness
// of errorflow usage and panics with message "errflow incorrect call sequence" when the issue
// is detected. Stacktrace will point to the location of the issue.
//
// Validation automatically checks for these rules:
//  * If any of Check* functions is used in a function, IfError() handler should be set up for this function.
//  * If IfError() handler is set up, it should be terminated by one of Then* functions.
//
// By default, validation is only enabled in tests and is disabled in production binaries.
//
// To enable/disable validation (e.g. because of performance issues, false positives,
// incorrect test/prod environment detection), clients can use SetNoopValidator
// or SetStackTraceValidator functions.
//
//  func Test_WithDisabledValidator(t *testing.T) {
//  	defer errf.SetNoopValidator().ThenRestore()
//
//  	// Validator will be disabled until the end of this function.
//  }
//
// Return Strategy
//
// Return strategy controls what error to return from a function in case if multiple errors
// were produced (e.g. by gzipWriter.Close() function and then fileWriter.Close() function).
//
// Return strategy is set by using one of Return* functions from IfError() handler.
//
// Available strategies:
//  * ReturnFirst() - return first encountered error.
//  * ReturnLast() - return last encountered error.
//  * ReturnWrapped() - return error where:
//                      * error message combines both messages;
//                      * first error instance is available via errors.Unwrap();
//                      * second error instance instance is discarded.
//  * ReturnCombined() - return error where:
//                      * error message combines both messages;
//                      * all error instances are available via errf.GetCombinedErrors() function.
//
// Example:
//  func example() (err error) {
//  	defer errf.IfError().Return*().ThenAssignTo(&err)
//  	defer errf.CheckErr(fmt.Errorf("error 3"))
//  	defer errf.CheckErr(fmt.Errorf("error 2"))
//  	defer errf.CheckErr(fmt.Errorf("error 1"))
//  	return nil
//  }
//
//  Results:
//  * ReturnFirst -> "error 1"
//  * ReturnLast -> "error 3"
//  * ReturnWrapped -> "error 1 (also: error 2) (also: error 3)"
//                     errors.Unwrap(err).Error() == "error 1"
//  * ReturnCombined -> "combined error {error 1; error 2; error 3}"
//                     errf.GetCombinedErrors(err) ->
//                       {fmt.Errorf("error 1"), fmt.Errorf("error 2"), fmt.Errorf("error 3")}
//
// Log Strategy
//
// Log strategy controls IfError() handler logging behavior.
//
// Available strategies:
//  * LogNever() - never log errors.
//  * LogIfSuppressed() - log only errors which are not included in result.
//  * LogAlways() - log all errors.
//
// Example:
//  func example() (err error) {
//  	defer errf.IfError().ReturnFirst().Log*().ThenAssignTo(&err)
//  	defer errf.CheckErr(fmt.Errorf("error 2"))
//  	defer errf.CheckErr(fmt.Errorf("error 1"))
//  	return nil
//  }
//
//  Results:
//  * LogNever -> no logs
//  * LogIfSuppressed -> "error 2" is logged
//  * LogAlways -> both "error 1" and "error 2" are logged
//
// Wrappers
//
// Wrappers are functions which wrap error objects into other error objects.
//
// They can be applied to:
//  * individual Check* functions
//  example: errf.With(wrapper).CheckErr(err)
//  * IfError() helper
//  example: defer errf.IfError().Apply(wrapper).ThenAssignTo(&err)
//  * helper functions (errf.Log, errf.IfErrorAssignTo)
//  example: defer errf.With(wrapper).Log(err)
//
// Example:
//
//  func example() (err error) {
//  	defer errf.IfError().Apply(errf.WrapperFmtErrorw("wrapped")).ThenAssignTo(&err)
//  	return errf.CheckErr(fmt.Errorf("error"))
//  }
//
//  Returns error with message: "wrapped: error"
//
// Custom wrappers can be implemented using errf.Wrapper function:
//
//  func addStacktraceToError(err error) error {
//  	/* ... actual implementation ... */
//  }
//
//  var WrapperAddStacktrace = func(ef *Errflow) *Errflow {
//  	return ef.With(errf.Wrapper(addStacktraceToError))
//  }
//
// Usage:
//
//  func example() (err error) {
//  	defer errf.IfError().Apply(WrapperAddStacktrace).ThenAssignTo(&err)
//
//  	// business logic ...
//  }
//
// Handlers
//
// Handlers are used to handle errors and panics that are bubbling up the defers ladder.
//
// Unlike IfError()... handler, Handle() API can be used in a middle of the function, thus
// e.g. skipped if preparation steps haven't been executed yet.
//
// This might be useful for additional cleanup logic. Consider function that
// copies files. If there is an error during copying, we would want to delete destination
// file rather than leaving it in inconsistent state. But, unlike closing a file, this action
// needs to be performed only when there's an error and should not be performed on a successful
// copy.
//
// Here's how handlers can be used to implement this logic:
//
//  func copyfile(src, dst string) (err error) {
//  	defer errf.IfError().ThenAssignTo(&err)
//
//  	/* some code */
//
//  	writer := errf.Io.WriteCloser(os.Create(filename))
//  	defer errf.Handle().OnAnyErrorOrPanic(func() { os.Remove(filename) })
//  	defer errf.CheckErr(writer.Close())
//
//  	/* more code */
//  }
//
// For full list of handlers, see documentation for:
//  * (*InterimHandler) On* methods
//  * (*InterimHandler) {Always, Everything}  methods
//
// Notes:
//  * Handle() API should be used only in defer statements.
//  * Handlers with "Err" in the name (e.g. OnErr, OnErrOrPanic) can only be used
//    in functions with IfError() handler.
//  * Handlers without "Err" in the name (e.g. Always, OnPanic) can be used
//    in any function.
//  * It is allowed to use Check* funcions inside Handlers even without IfError() set up
//    inside a handler. In such cases, defer Handle()... enclosing function IfError()
//    will be used to catch errors.
//
// Custom log function
//
// Custom log function can be set using SetLogFn method:
//
//  func customLogFn(logMessage *LogMessage) {
//  	/* ... implementation ... */
//  }
//
//  func main() {
//  	defer errf.SetLogFn(customLogFn).ThenRestore()
//
//  	/* app main function */
//  }
//
// Helper functions
//
// Errflow also implements few helper functions which can be used
// in functions without IfError() handler.
//
//  errf.Log(err) will log error, if not nil. It doesn't affect control flow.
//
//  defer errf.IfErrorThenAssign(&err, operationErr) will assign (*err = operationErr) if operationErr != nil.
//    Note: it is useful only in 'defer' for function without IfError() handler
//          as a lightweight alternative.
//
// This functions might be configured using errf.With(...).
//
// NOTE: since these functions don't use IfError() handler, they
// will not use config defined on IfError() handler.
//
//  func example() (err error) {
//  	errOpts = errf.WrapperFmtErrorw("error in example function")
//  	defer errf.IfError().Apply(errOpts).ThenAssignTo(&err)
//
//  	errf.With(errOpts).Log(fmt.Errorf("operation 1 error"))
//  	return errf.CheckErr(fmt.Errorf("operation 2 error"))
//  }
//
//  This function will return error "error in example function: operation 2 error".
//  Also it will log "error in example function: operation 1 error".
//  Note that Log requires .With(errOpts), because it won't use IfError() context.
//
// Extending errorflow for custom types
//
// Library clients can create new Check* functions for custom return types.
// This is useful for writing type safe code (until generics will solve this problem).
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
//  	customStructValue := fancypackage.Errf.CheckCustomType1(fancypackage.ReadCustomType1())
//
//  	customStructValue := fancypackage.Errf.
//  		With(errf.WrapperFmtErrorw("error in ProcessCustomStruct")).
//  		CheckCustomType2(fancypackage.ReadCustomType2())
//
//  	// ...
//  }
package errf
