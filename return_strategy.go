package errf

import (
	"errors"
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

func getReturnStrategyImpl(rs returnStrategy) func(error, error) (supp1, supp2 bool, result error) {
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

// ReturnStrategyFirst configures Errflow instance to return first error instance.
// This is default behavior.
func ReturnStrategyFirst(ef *Errflow) *Errflow {
	return setReturnStrategy(ef, returnStrategyFirst)
}

// ReturnStrategyLast configures Errflow instance to return last error instance.
func ReturnStrategyLast(ef *Errflow) *Errflow {
	return setReturnStrategy(ef, returnStrategyLast)
}

// ReturnStrategyWrapped configures Errflow instance to return all errors.
// First error will be wrapped using fmt.Errorf with "%w" parameter.
// For other error, their messages will be included in resulting errors,
// but the instances will be discarded.
func ReturnStrategyWrapped(ef *Errflow) *Errflow {
	return setReturnStrategy(ef, returnStrategyWrapped)
}

// ReturnStrategyCombined configures Errflow instance to return all errors.
// All error messages will be collected in resulting message
// and all error instances will be preserved.
// Error instances can be retrieved using GetCombinedErrors() function.
func ReturnStrategyCombined(ef *Errflow) *Errflow {
	return setReturnStrategy(ef, returnStrategyCombined)
}

func returnStrategyFirstImpl(err1, _ error) (supp1, supp2 bool, result error) {
	return false, true, err1
}

func returnStrategyLastImpl(_, err2 error) (supp1, supp2 bool, result error) {
	return true, false, err2
}

func returnStrategyWrappedImpl(err1, err2 error) (supp1, supp2 bool, result error) {
	return false, false, fmt.Errorf("%w (also: %s)", err1, err2.Error())
}

// CombinedError implements an error that holds multiple errors.
type CombinedError struct {
	errs []error
}

func (cErr CombinedError) Error() string {
	var errorList []string
	for _, err := range cErr.errs {
		errorList = append(errorList, err.Error())
	}
	return fmt.Sprintf("combined error {%s}", strings.Join(errorList, "; "))
}

// GetCombinedErrors returns all error instances from CombinedError error,
// even if the error was wrapped using fmt.Errorf with "%w" parameter.
//
// Note that resulting errors are all flattened out into a single list,
// meaning that calling GetCombinedErrors on errors returned from GetCombinedErrors
// will always result in returning the same error.
func GetCombinedErrors(err error) []error {
	if err == nil {
		return nil
	}
	var cErr CombinedError
	if errors.As(err, &cErr) {
		return cErr.errs
	}
	return []error{err}
}

func returnStrategyCombinedImpl(err1, err2 error) (supp1, supp2 bool, result error) {
	var errs []error

	errs = append(errs, GetCombinedErrors(err1)...)
	errs = append(errs, GetCombinedErrors(err2)...)

	if len(errs) == 0 {
		return false, false, nil
	} else if len(errs) == 1 {
		return false, false, errs[0]
	} else {
		return false, false, CombinedError{
			errs: errs,
		}
	}
}
