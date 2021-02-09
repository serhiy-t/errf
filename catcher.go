package errflow

// Catcher controls error handling behavior.
// See IfError() for more info.
type Catcher interface {

	// ThenAssignTo is a teminal statement, which instructs to write error value
	// to 'outErr', in case there is an error.
	ThenAssignTo(outErr *error)

	// Then is a teminal statement, which instructs to call a callback function
	// with error value, in case there is an error.
	Then(fn func(err error))

	// ThenIgnore is a teminal statement, which instructs to ignore errors encountered
	// in this function.
	ThenIgnore()

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

// IfError creates a Catcher instance to control error handling behavior.
//
// Example:
//  func function() (err error) {
//    defer errflow.IfError().WriteTo(&err)
//
//    // Function definition.
//  }
//
// Usage notes:
//  * errflow.IfError() should be called in a defer, as a first statement.
//  * errflow.IfError() should be always terminated by one of .Then*(...) functions.
//  * it should only process errors returned by the same function where it is declared;
//    otherwise validation will fail during when running tests.
func IfError() Catcher {
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

func (c *catcher) ThenAssignTo(outErr *error) {
	c.catch(recover(), func(err error) { *outErr = err })
}

func (c *catcher) Then(fn func(err error)) {
	c.catch(recover(), fn)
}

func (c *catcher) ThenIgnore() {
	c.catch(recover(), func(err error) {})
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

func (c *catcher) catch(recoverObj interface{}, fn func(err error)) {
	globalErrflowValidator.leave()
	if recoverObj != nil {
		errflowThrow, ok := recoverObj.(errflowThrow)
		if ok {
			if c.loggerFn != nil {
				for _, err := range errflowThrow.errs {
					c.loggerFn(err.Error())
				}
			}
			fn(c.returnErrorStrategy.returnError(errflowThrow.errs))
		} else {
			panic(recoverObj)
		}
	}
}
