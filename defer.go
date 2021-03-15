package errf

import (
	"fmt"
)

func (ef *Errflow) IfErrorAssignTo(outErr *error, err error) {
	ef.applyDeferredOptions()
	if maySuppressFirstError(ef.returnStrategy) {
		panic(fmt.Errorf("%v is not supported for IfErrorAssignTo(...)", ef.returnStrategy))
	}
	if err != nil {
		if ef.wrapper != nil {
			err = ef.wrapper(err)
		}
		if *outErr == nil {
			*outErr = err
			if ef.logStrategy == logStrategyAlways {
				globalLogFn(&LogMessage{
					Format: "%s",
					A:      []interface{}{err.Error()},
					Stack:  getStringErrorStackTraceFn(),
					Tags:   []string{"errflow", "error"},
				})
			}
		} else {
			resultErr, _, supp2 := getReturnStrategyImpl(ef.returnStrategy)(*outErr, err)
			*outErr = resultErr
			if (supp2 && ef.logStrategy == logStrategyIfSuppressed) || ef.logStrategy == logStrategyAlways {
				globalLogFn(&LogMessage{
					Format: "%s",
					A:      []interface{}{err.Error()},
					Stack:  getStringErrorStackTraceFn(),
					Tags:   []string{"errflow", "suppressed-error"},
				})
			}
		}
	}
}

func IfErrorAssignTo(outErr *error, err error) {
	DefaultErrflow.IfErrorAssignTo(outErr, err)
}
