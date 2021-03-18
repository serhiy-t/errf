package errf

import (
	"errors"
	"fmt"
	"reflect"
)

type handleCondition struct {
	onError     bool
	onSuccess   bool
	onPanic     bool
	notValidate bool
}

// InterimHandler defines Handle() API.
type InterimHandler struct{}

// Handle enables additional error handlers in the middle of functions,
// in addition to IfError() handlers.
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
// Example:
//  func example(filename string) (err error) {
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
func Handle() *InterimHandler {
	return &InterimHandler{}
}

// PanicErr is an error type, which is used in error fn Handle()... callbacks, in case if handler
// was triggered by a panic instead of an error.
//
// See also: errf.IsPanic, errf.GetPanic.
type PanicErr struct {
	PanicObj interface{}
}

func (p PanicErr) Error() string {
	return fmt.Sprintf("panic: %v", p.PanicObj)
}

// Always handler is always executed.
// Error is not sent to the callback.
//
// Use Everything(), if error info is required.
func (h *InterimHandler) Always(errFn func()) {
	h.handle(recover(), handleCondition{onError: true, onPanic: true, onSuccess: true, notValidate: true}, func(err error) {
		errFn()
	})
}

// Everything handler is always executed.
// Errors and panics are sent to the callback.
//
// Use IsPanic(), IsErr(), IsSuccess() to differentiate between those outcomes.
//
// Use Always(), if error info is not needed.
func (h *InterimHandler) Everything(errFn ErrorActionFn) {
	h.handle(recover(), handleCondition{onError: true, onPanic: true, onSuccess: true, notValidate: true}, errFn)
}

// OnErr handler is executed in case of error triggered by one of "Check*" functions.
//
// First encountered error is passed to the callback.
func (h *InterimHandler) OnErr(errFn ErrorActionFn) {
	h.handle(recover(), handleCondition{onError: true}, errFn)
}

// OnErrIs handler is executed in case of error triggered by one of "Check*" functions
// and first encountered error is targetErr (using errors.Is definition).
func (h *InterimHandler) OnErrIs(targetErr error, errFn func()) {
	h.handle(recover(), handleCondition{onError: true}, func(err error) {
		if errors.Is(err, targetErr) {
			errFn()
		}
	})
}

func verifyErrFnType(argument string, errFn interface{}) {
	t := reflect.TypeOf(errFn)
	if t.Kind() != reflect.Func {
		panic(fmt.Errorf("%s should be a function", argument))
	}
	if t.NumIn() != 1 || t.IsVariadic() {
		panic(fmt.Errorf("%s should have exactly 1 input argument", argument))
	}
	if t.NumOut() != 0 {
		panic(fmt.Errorf("%s should have exactly no output arguments", argument))
	}
	errType := reflect.TypeOf((*error)(nil)).Elem()
	if !t.In(0).AssignableTo(errType) {
		panic(fmt.Errorf("%s first argument should be assignable to error interface", argument))
	}
}

// OnErrAs handler is executed in case of error triggered by one of "Check*" functions
// and first encountered error has type of callback argument (using errors.As definition).
//
// Example:
//  defer errf.Handle().OnErrAs(func (err net.Error) {
//  	// This callback only will be executed if first encountered
//  	// error has type of net.Error.
//  })
func (h *InterimHandler) OnErrAs(errFn interface{}) {
	globalErrflowValidator.custom(func() {
		verifyErrFnType("OnErrAs: errFn", errFn)
	})
	h.handle(recover(), handleCondition{onError: true}, func(err error) {
		errFnValue := reflect.ValueOf(errFn)
		errValue := reflect.New(errFnValue.Type().In(0))
		if errors.As(err, errValue.Interface()) {
			errFnValue.Call([]reflect.Value{errValue.Elem()})
		}
	})
}

// OnErrOrPanic handler is executed in case of error triggered by one of "Check*" functions
// or a panic.
//
// First encountered error is passed to the callback.
// See also errf.IsPanic(), errf.IsErr().
func (h *InterimHandler) OnErrOrPanic(errFn ErrorActionFn) {
	h.handle(recover(), handleCondition{onError: true, onPanic: true}, errFn)
}

// OnPanic handler is executed in case of a panic.
func (h *InterimHandler) OnPanic(panicFn func(panicObj interface{})) {
	h.handle(recover(), handleCondition{onPanic: true}, func(err error) {
		panicFn(err.(PanicErr).PanicObj)
	})
}

// OnAnyPanic handler is same as OnPanic, when panicObj is not required.
func (h *InterimHandler) OnAnyPanic(panicFn func()) {
	h.handle(recover(), handleCondition{onPanic: true}, func(err error) {
		panicFn()
	})
}

// OnAnyErr handler is same as OnErr, when err is not required.
func (h *InterimHandler) OnAnyErr(errFn func()) {
	h.handle(recover(), handleCondition{onError: true}, func(err error) { errFn() })
}

// OnAnyErrOrPanic handler is same as OnErrOrPanic, when err is not required.
func (h *InterimHandler) OnAnyErrOrPanic(errFn func()) {
	h.handle(recover(), handleCondition{onError: true, onPanic: true}, func(err error) { errFn() })
}

// OnSuccess handler is executed in case of no errors or panics.
func (h *InterimHandler) OnSuccess(successFn func()) {
	h.handle(recover(), handleCondition{onSuccess: true}, func(err error) { successFn() })
}

// IsErr returns true when error send to handler callback indicates an error (not panic or success).
// Useful for handlers which handle multiple types (e.g. Everything(), OnErrOrPanic())
func IsErr(err error) bool {
	return err != nil && !IsPanic(err)
}

// IsSuccess returns true when error send to handler callback indicates success
// (is null, no errors or panics).
// Useful for handlers which handle multiple types (e.g. Everything())
func IsSuccess(err error) bool {
	return err == nil
}

// IsPanic returns true when error send to handler callback indicates a panic (not error or success).
// Useful for handlers which handle multiple types (e.g. Everything(), OnErrOrPanic())
func IsPanic(err error) bool {
	_, ok := err.(PanicErr)
	return ok
}

// GetPanic returns true when error send to handler callback indicates a panic (not error or success).
// Also it writes panic value into panicObj pointer.
func GetPanic(err error, panicObj *interface{}) bool {
	panicErr, ok := err.(PanicErr)
	if ok {
		*panicObj = panicErr.PanicObj
	}
	return ok
}

func (h *InterimHandler) handle(
	recoverObj interface{},
	condition handleCondition,
	fn ErrorActionFn,
) {
	if condition.onError && !condition.notValidate {
		if isUnrelatedPanic(recoverObj) {
			globalErrflowValidator.markPanic()
		}

		globalErrflowValidator.validate()
	}

	if recoverObj != nil {
		errflowThrowObj, ok := recoverObj.(errflowThrow)
		if ok && len(errflowThrowObj.items) > 0 {
			item := errflowThrowObj.items[0]
			ef := item.ef
			err := item.err
			ef.applyDeferredOptions()
			if ef.wrapper != nil && err != nil {
				err = ef.wrapper(err)
			}
			defer func() {
				fnRecover := recover()
				fnErrflowThrow, isFnErrflowThrow := fnRecover.(errflowThrow)
				if fnRecover == nil {
					panic(errflowThrowObj)
				} else if isFnErrflowThrow {
					var combinedErrflowThrow errflowThrow
					combinedErrflowThrow.items = append(combinedErrflowThrow.items, errflowThrowObj.items...)
					combinedErrflowThrow.items = append(combinedErrflowThrow.items, fnErrflowThrow.items...)
					panic(combinedErrflowThrow)
				} else {
					panic(fnRecover)
				}
			}()
			if condition.onError {
				fn(err)
			}
		} else {
			defer func() {
				fnRecover := recover()
				_, isFnErrflowThrow := fnRecover.(errflowThrow)
				if fnRecover == nil {
					panic(recoverObj)
				} else if isFnErrflowThrow {
					panic(recoverObj)
				} else {
					panic(fnRecover)
				}
			}()
			if condition.onPanic {
				fn(PanicErr{PanicObj: recoverObj})
			}
		}
	} else {
		if condition.onSuccess {
			fn(nil)
		}
	}
}
