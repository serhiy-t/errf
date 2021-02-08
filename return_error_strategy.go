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

func GetAllErrors(err error) []error {
	all, ok := err.(*allErrors)
	if ok {
		return all.errs
	} else if err != nil {
		return []error{err}
	} else {
		return nil
	}
}
