package errflow

import (
	"fmt"
	"runtime/debug"
)

// Catcher controls error handling behavior.
// See Catch() for more info.
type Catcher interface {

	// WriteTo is a teminal statement, which instructs to write error value
	// to 'outErr', in case there is an error.
	WriteTo(outErr *error)

	// Then is a teminal statement, which instructs to call a callback function
	// with error value, in case there is an error.
	Then(fn func(err error))

	// ReturnFirst is a non-terminal statement which configures Catcher to
	// return a first encountered error in case of multiple errors.
	// This is a default behavior.
	ReturnFirst() Catcher

	// ReturnLast is a non-terminal statement which configures Catcher to
	// return a last encountered error in case of multiple errors.
	ReturnLast() Catcher
	// ReturnAll is a non-terminal statement which configures Catcher to

	// return all errors in case of multiple errors.
	// See also errflow.GetAllErrors(...).
	ReturnAll() Catcher

	// LogAll is a non-terminal statement which configures Catcher to
	// log all errors using logger function set via errflow.SetLogFn(...).
	LogAll() Catcher

	// LogNone is a non-terminal statement which configures Catcher to
	// not not errors.
	// This is a default behavior.
	LogNone() Catcher
}

// Catch creates a Catcher instance to control error handling behavior.
//
// Example:
//  func function() (err error) {
//    defer errflow.Catch().WriteTo(&err)
//
//    // Function definition.
//  }
//
// Usage notes:
//  * errflow.Catch() should be called in a defer, as a first statement.
//  * errflow.Catch() should be always terminated by .WriteTo(...) or .Then(...).
//  * it should only process errors returned by the same function where it is declared;
//    otherwise validation will fail during when running tests.
func Catch() Catcher {
	globalErrflowValidator.enter()
	return &catcher{
		returnErrorStrategy: &returnErrorStrategyFirst{},
		loggerFn:            nil,
	}
}

type catcher struct {
	returnErrorStrategy returnErrorStrategy
	loggerFn            func(s string)
}

func (c *catcher) WriteTo(outErr *error) {
	c.catch(recover(), func(err error) { *outErr = err })
}

func (c *catcher) Then(fn func(err error)) {
	c.catch(recover(), fn)
}

func (c *catcher) ReturnFirst() Catcher {
	c.returnErrorStrategy = &returnErrorStrategyFirst{}
	return c
}

func (c *catcher) ReturnLast() Catcher {
	c.returnErrorStrategy = &returnErrorStrategyLast{}
	return c
}

func (c *catcher) ReturnAll() Catcher {
	c.returnErrorStrategy = &returnErrorStrategyAll{}
	return c
}

func (c *catcher) LogAll() Catcher {
	c.loggerFn = globalLogFn
	return c
}

func (c *catcher) LogNone() Catcher {
	c.loggerFn = nil
	return c
}

func (c *catcher) catch(panicObj interface{}, fn func(err error)) {
	globalErrflowValidator.leave()
	if panicObj != nil {
		errflowThrow, ok := panicObj.(errflowThrow)
		if ok {
			if c.loggerFn != nil {
				for _, err := range errflowThrow.errs {
					c.loggerFn(err.Error())
				}
			}
			fn(c.returnErrorStrategy.returnError(errflowThrow.errs))
		} else {
			c.loggerFn(fmt.Sprintf("Panic: %s\n%s", panicObj, debug.Stack()))
			panic(panicObj)
		}
	}
}
