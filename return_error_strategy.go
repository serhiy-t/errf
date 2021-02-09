package errflow

import "fmt"

type returnErrorStrategy interface {
	returnError(errs []error) error
}

type returnErrorStrategyFirst struct{}

func (s *returnErrorStrategyFirst) returnError(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	return errs[0]
}

type returnErrorStrategyLast struct{}

func (s *returnErrorStrategyLast) returnError(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	return errs[len(errs)-1]
}

type allErrors struct {
	errs []error
}

func (ae *allErrors) Error() string {
	return fmt.Sprintf("Multiple errors: %s", ae.errs)
}

type returnErrorStrategyAll struct{}

func (s *returnErrorStrategyAll) returnError(errs []error) error {
	if len(errs) == 0 {
		return nil
	} else if len(errs) == 1 {
		return errs[0]
	}
	return &allErrors{errs: errs}
}

// GetAllErrors unpacks all errors into a slice when error is produced
// using ReturnAll() strategy.
func GetAllErrors(err error) []error {
	return getAllErrorsInternal(err, false)
}

// GetAllErrorsFlattened unpacks all errors into a slice when error is produced
// using ReturnAll() strategy.
// Result is flattened, meaning if any error containes nested multiple errors, all will
// be unpacked.
func GetAllErrorsFlattened(err error) []error {
	return getAllErrorsInternal(err, true)
}

func getAllErrorsInternal(err error, flattened bool) []error {
	var result []error

	all, ok := err.(*allErrors)
	if ok {
		for _, nestedErr := range all.errs {
			if flattened {
				result = append(result, getAllErrorsInternal(nestedErr, true)...)
			} else {
				result = append(result, nestedErr)
			}
		}
	} else if err != nil {
		result = append(result, err)
	}

	return result
}
