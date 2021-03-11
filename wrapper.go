package errflow

import "fmt"

func Wrapper(wrapper func(err error) error) ErrflowOption {
	return func(ef *Errflow) *Errflow {
		newEf := ef.Copy()
		newEf.wrapper = func(err error) error {
			return wrapper(ef.wrapper(err))
		}
		return newEf
	}
}

func fmtErrorf(format string, a ...interface{}) func(err error) error {
	return func(err error) error {
		for i, v := range a {
			if v == OriginalErr {
				a[i] = err
			}
		}
		return fmt.Errorf(format, a...)
	}
}

func WrapperFmtErrorf(format string, a ...interface{}) ErrflowOption {
	return Wrapper(fmtErrorf(format, a...))
}

func WrapperFmtErrorW(s string) ErrflowOption {
	return WrapperFmtErrorf("%s: %w", s, OriginalErr)
}

type originalErrType struct{}

func (err *originalErrType) Error() string {
	return "errflow original error placeholder"
}

var OriginalErr = &originalErrType{}
