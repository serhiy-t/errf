package errflow

func IfError() *IfErrorHandler {
	globalErrflowValidator.enter()
	return &IfErrorHandler{}
}

type IfErrorHandler struct {
	options []ErrflowOption
}

func (c *IfErrorHandler) ThenAssignTo(outErr *error) {
	c.catch(recover(), func(err error) { *outErr = err })
}

func (c *IfErrorHandler) Then(fn func(err error)) {
	c.catch(recover(), fn)
}

func (c *IfErrorHandler) ThenIgnore() {
	c.catch(recover(), func(err error) {})
}

func (c *IfErrorHandler) ReturnFirst() *IfErrorHandler {
	return c.Apply(ReturnStrategyFirst)
}

func (c *IfErrorHandler) ReturnLast() *IfErrorHandler {
	return c.Apply(ReturnStrategyLast)
}

func (c *IfErrorHandler) ReturnWrapped() *IfErrorHandler {
	return c.Apply(ReturnStrategyWrapped)
}

func (c *IfErrorHandler) ReturnCombined() *IfErrorHandler {
	return c.Apply(ReturnStrategyCombined)
}

func (c *IfErrorHandler) LogAlways() *IfErrorHandler {
	return c.Apply(LogStrategyAlways)
}

func (c *IfErrorHandler) LogIfSuppressed() *IfErrorHandler {
	return c.Apply(LogStrategyIfSuppressed)
}

func (c *IfErrorHandler) LogNever() *IfErrorHandler {
	return c.Apply(LogStrategyNever)
}

func (c *IfErrorHandler) Apply(options ...ErrflowOption) *IfErrorHandler {
	c.options = append(c.options, options...)
	return c
}

func (c *IfErrorHandler) catch(recoverObj interface{}, fn func(err error)) {
	globalErrflowValidator.leave()

	if recoverObj != nil {
		errflowThrow, ok := recoverObj.(errflowThrow)
		if ok {
			var currItem errflowThrowItem
			for _, item := range errflowThrow.items {
				item.ef = item.ef.With(c.options...)
				item.ef.applyDeferredOptions()

				if item.ef.logStrategy == logStrategyAlways {
					globalLogFn(&LogMessage{
						Format: "%s",
						A:      []interface{}{item.err.Error()},
						Stack:  getErrorStackTrace(),
						Tags:   []string{"errflow", "error"},
					})
				}

				if !(currItem.ef == nil && currItem.err == nil) {
					newErr, supp1, supp2 := getReturnStrategyImpl(item.ef.returnStrategy)(currItem.err, item.err)

					if supp1 && currItem.ef.logStrategy == logStrategyIfSuppressed {
						globalLogFn(&LogMessage{
							Format: "%s",
							A:      []interface{}{currItem.err.Error()},
							Stack:  getErrorStackTrace(),
							Tags:   []string{"errflow", "suppressed-error"},
						})
					}
					if supp2 && item.ef.logStrategy == logStrategyIfSuppressed {
						globalLogFn(&LogMessage{
							Format: "%s",
							A:      []interface{}{item.err.Error()},
							Stack:  getErrorStackTrace(),
							Tags:   []string{"errflow", "suppressed-error"},
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
