package errf

// IfError creates IfErrorHandler.
//
// Should always:
//   * be used only in defer statements;
//   * be in the beginning of a function;
//   * terminated by one of Then*(...) functions.
//
// Example:
//  function example() (err error) {
//    defer errf.IfError().ThenAssignTo(&err)
//
//    // ...
//  }
func IfError() *IfErrorHandler {
	globalErrflowValidator.enter()
	return &IfErrorHandler{}
}

// IfErrorHandler configures errorflow error handling behavior in a scope of a function.
//
// Should be created only via IfError() method.
type IfErrorHandler struct {
	options []ErrflowOption
}

// ThenAssignTo assigns resulting error to outErr (only if non-nil).
// If outErr is already non-nil, it will be replaced by error returned
// from IfErrorHandler, and logged if log strategy is IfSuppressed or Always.
//
// Note: it is not recommended to mix returning errors directly and via Check* function
// because IfError handler doesn't have much control over direct errors,
// which might result in counterintuitive behavior.
//
// To avoid mixing, always instead of:
//   return err
// write
//   return errf.CheckErr(err).IfOkReturnNil
func (c *IfErrorHandler) ThenAssignTo(outErr *error) {
	c.catch(recover(), func(err error) {
		if *outErr != nil {
			ef := With(c.options...)
			ef.applyDeferredOptions()
			if ef.logStrategy == logStrategyAlways || ef.logStrategy == logStrategyIfSuppressed {
				globalLogFn(&LogMessage{
					Format: "%s",
					A:      []interface{}{(*outErr).Error()},
					Stack:  getStringErrorStackTraceFn(),
					Tags:   []string{"errorflow", "suppressed-external-error"},
				})
			}
		}
		*outErr = err
	})
}

// ErrorActionFn is a callback for errors handlers.
type ErrorActionFn func(err error)

// Then calls callbacks for resulting error (only if non-nil).
func (c *IfErrorHandler) Then(fns ...ErrorActionFn) {
	c.catch(recover(), func(err error) {
		for _, fn := range fns {
			fn(err)
		}
	})
}

// ThenIgnore ignores resulting error.
func (c *IfErrorHandler) ThenIgnore() {
	c.catch(recover(), func(err error) {})
}

// ReturnFirst is an alias for Apply(ReturnStrategyFirst).
func (c *IfErrorHandler) ReturnFirst() *IfErrorHandler {
	return c.Apply(ReturnStrategyFirst)
}

// ReturnLast is an alias for Apply(ReturnStrategyLast).
func (c *IfErrorHandler) ReturnLast() *IfErrorHandler {
	return c.Apply(ReturnStrategyLast)
}

// ReturnWrapped is an alias for Apply(ReturnStrategyWrapped).
func (c *IfErrorHandler) ReturnWrapped() *IfErrorHandler {
	return c.Apply(ReturnStrategyWrapped)
}

// ReturnCombined is an alias for Apply(ReturnStrategyCombined).
func (c *IfErrorHandler) ReturnCombined() *IfErrorHandler {
	return c.Apply(ReturnStrategyCombined)
}

// LogAlways is an alias for Apply(LogStrategyAlways).
func (c *IfErrorHandler) LogAlways() *IfErrorHandler {
	return c.Apply(LogStrategyAlways)
}

// LogIfSuppressed is an alias for Apply(LogStrategyIfSuppressed).
func (c *IfErrorHandler) LogIfSuppressed() *IfErrorHandler {
	return c.Apply(LogStrategyIfSuppressed)
}

// LogNever is an alias for Apply(LogStrategyNever).
func (c *IfErrorHandler) LogNever() *IfErrorHandler {
	return c.Apply(LogStrategyNever)
}

// Apply configures IfErrorHandler to apply additional configs to Errflow.
//
// Example:
//  func example() (err error) {
//  	defer errf.IfError().Apply(errf.WrapperFmtErrorw("wrapper 2")).ThenAssignTo(&err)
//
//  	With(errf.WrapperFmtErrorw("wrapper 1")).CheckErr(fmt.Errorf("error 1"))
//  	return nil
//  }
//
// When example() is called, "wrapper 1" will be applied first and "wrapper 2" will be applied second.
//
// Resulting error message is: "wrapper 2: wrapper 1: error 1".
func (c *IfErrorHandler) Apply(options ...ErrflowOption) *IfErrorHandler {
	c.options = append(c.options, options...)
	return c
}

func isUnrelatedPanic(recoverObj interface{}) bool {
	if recoverObj != nil {
		_, ok := recoverObj.(errflowThrow)
		if !ok {
			return true
		}
	}
	return false
}

func (c *IfErrorHandler) catch(recoverObj interface{}, fn ErrorActionFn) {
	if isUnrelatedPanic(recoverObj) {
		globalErrflowValidator.markPanic()
	}
	globalErrflowValidator.leave()

	if recoverObj != nil {
		errflowThrow, ok := recoverObj.(errflowThrow)
		if ok {
			var currItem errflowThrowItem
			for _, item := range errflowThrow.items {
				item.ef = item.ef.With(c.options...)
				item.ef.applyDeferredOptions()
				if item.ef.wrapper != nil && item.err != nil {
					item.err = item.ef.wrapper(item.err)
				}

				if item.ef.logStrategy == logStrategyAlways {
					globalLogFn(&LogMessage{
						Format: "%s",
						A:      []interface{}{item.err.Error()},
						Stack:  getStringErrorStackTraceFn(),
						Tags:   []string{"errorflow", "error"},
					})
				}

				if !(currItem.ef == nil && currItem.err == nil) {
					supp1, supp2, newErr := getReturnStrategyImpl(item.ef.returnStrategy)(currItem.err, item.err)

					if supp1 && currItem.ef.logStrategy == logStrategyIfSuppressed {
						globalLogFn(&LogMessage{
							Format: "%s",
							A:      []interface{}{currItem.err.Error()},
							Stack:  getStringErrorStackTraceFn(),
							Tags:   []string{"errorflow", "suppressed-error"},
						})
					}
					if supp2 && item.ef.logStrategy == logStrategyIfSuppressed {
						globalLogFn(&LogMessage{
							Format: "%s",
							A:      []interface{}{item.err.Error()},
							Stack:  getStringErrorStackTraceFn(),
							Tags:   []string{"errorflow", "suppressed-error"},
						})
					}

					currItem.err = newErr
					currItem.ef = item.ef
				} else {
					currItem = item
				}
			}
			fn(currItem.err)
		} else {
			panic(recoverObj)
		}
	}
}
