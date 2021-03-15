package errf

import "fmt"

func Wrapper(wrapper func(err error) error) ErrflowOption {
	return func(ef *Errflow) *Errflow {
		if wrapper == nil {
			return ef
		}
		newEf := ef.copy()
		oldWrapper := ef.wrapper
		newWrapper := wrapper

		if oldWrapper != nil {
			newWrapper = func(err error) error {
				return wrapper(oldWrapper(err))
			}
		}

		newEf.wrapper = newWrapper
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

func WrapperFmtErrorw(s string) ErrflowOption {
	return WrapperFmtErrorf("%s: %w", s, OriginalErr)
}

type originalErrType struct{}

func (err *originalErrType) Error() string {
	return "errflow original error placeholder"
}

var OriginalErr = &originalErrType{}
