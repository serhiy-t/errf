package errf

import (
	"fmt"
	"strings"
)

type returnStrategy int

const (
	returnStrategyDefault = returnStrategy(iota)
	returnStrategyUnknown
	returnStrategyFirst
	returnStrategyLast
	returnStrategyWrapped
	returnStrategyCombined
)

func maySuppressFirstError(rs returnStrategy) bool {
	return rs == returnStrategyLast
}

func getReturnStrategyImpl(rs returnStrategy) func(error, error) (result error, supp1, supp2 bool) {
	switch rs {
	case returnStrategyFirst:
		return returnStrategyFirstImpl
	case returnStrategyLast:
		return returnStrategyLastImpl
	case returnStrategyWrapped:
		return returnStrategyWrappedImpl
	case returnStrategyCombined:
		return returnStrategyCombinedImpl
	}

	if rs != returnStrategyDefault {
		panic(fmt.Errorf("unknown errflow return strategy: %v", rs))
	}
	return returnStrategyFirstImpl
}

func setReturnStrategy(ef *Errflow, rs returnStrategy) *Errflow {
	newEf := ef.copy()
	if ef.returnStrategy == returnStrategyDefault {
		newEf.returnStrategy = rs
	}
	return newEf
}

func ReturnStrategyFirst(ef *Errflow) *Errflow {
	return setReturnStrategy(ef, returnStrategyFirst)
}

func ReturnStrategyLast(ef *Errflow) *Errflow {
	return setReturnStrategy(ef, returnStrategyLast)
}

func ReturnStrategyWrapped(ef *Errflow) *Errflow {
	return setReturnStrategy(ef, returnStrategyWrapped)
}

func ReturnStrategyCombined(ef *Errflow) *Errflow {
	return setReturnStrategy(ef, returnStrategyCombined)
}

func returnStrategyFirstImpl(err1, err2 error) (result error, supp1, supp2 bool) {
	return err1, false, true
}

func returnStrategyLastImpl(err1, err2 error) (result error, supp1, supp2 bool) {
	return err2, true, false
}

func returnStrategyWrappedImpl(err1, err2 error) (result error, supp1, supp2 bool) {
	return fmt.Errorf("%w (also: %s)", err1, err2.Error()), false, false
}

type CombinedError struct {
	errs []error
}

func (cerr CombinedError) Error() string {
	var errors []string
	for _, err := range cerr.errs {
		errors = append(errors, err.Error())
	}
	return fmt.Sprintf("combined error {%s}", strings.Join(errors, "; "))
}

func GetCombinedErrors(err error) []error {
	if err == nil {
		return nil
	}
	cerr, ok := err.(CombinedError)
	if ok {
		return cerr.errs
	}
	return []error{err}
}

func returnStrategyCombinedImpl(err1, err2 error) (result error, supp1, supp2 bool) {
	var errs []error

	errs = append(errs, GetCombinedErrors(err1)...)
	errs = append(errs, GetCombinedErrors(err2)...)

	if len(errs) == 0 {
		return nil, false, false
	} else if len(errs) == 1 {
		return errs[0], false, false
	} else {
		return CombinedError{
			errs: errs,
		}, false, false
	}
}
