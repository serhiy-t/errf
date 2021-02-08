package errflow

import (
	"io"
	"log"
)

// Catcher ...
type Catcher interface {
	WriteTo(outErr *error)
	Then(fn func(err error))
	ReturnFirst() Catcher
	ReturnLast() Catcher
	ReturnAll() Catcher
	LogAll() Catcher
	LogNone() Catcher
}

// Catch ...
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

var globalLogFn = func(s string) { log.Print(s) }

type restoreLogFn struct {
	oldLogFn func(s string)
}

func (rlf *restoreLogFn) Close() error {
	globalLogFn = rlf.oldLogFn
	return nil
}

func SetLogFn(logFn func(s string)) io.Closer {
	oldLogFn := globalLogFn
	globalLogFn = logFn
	return &restoreLogFn{
		oldLogFn: oldLogFn,
	}
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
			panic(panicObj)
		}
	}
}
