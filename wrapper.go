package errf

import "fmt"

// Wrapper function creates ErrflowOption that wraps original errors
// using provided 'func(err error) error' function.
// See WrapperFmtErrorw for common scenario fmt.Errorf("wrap message: %w", err).
//
// Example:
//  func WrapInOurError(ef *Errflow) *Errflow {
//  	return ef.With(errf.Wrapper(func (err error) error {
//  		return ourerrors.Wrap(err)
//  	}))
//  }
//
//  func exampleUsage() (err error) {
//  	defer errf.IfError().Apply(WrapInOurError).ThenAssignTo(&err)
//  	// ...
//  }
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

// WrapperFmtErrorf is a Wrapper that uses fmt.Errorf to wrap errors.
// errf.OriginalErr is used as a placeholder for error in fmt.Errorf args.
//
// See WrapperFmtErrorw for common scenario fmt.Errorf("wrap message: %w", err).
//
// Example:
//  func exampleUsage() (err error) {
//  	defer errf.IfError().Apply(
//  		errf.WrapperFmtErrorf("--> %w <--", errf.OriginalErr)
//  	).ThenAssignTo(&err)
//
//  	// ...
//  }
func WrapperFmtErrorf(format string, a ...interface{}) ErrflowOption {
	return Wrapper(fmtErrorf(format, a...))
}

// WrapperFmtErrorw is a Wrapper that uses fmt.Errorf("%s: %w", ...) to wrap errors.
//
// See WrapperFmtErrorw for common scenario fmt.Errorf("wrap message: %w", err).
//
// Example:
//  func exampleUsage() (err error) {
//  	defer errf.IfError().Apply(errf.WrapperFmtErrorw("error in exampleUsage")).ThenAssignTo(&err)
//
//  	// ...
//  }
func WrapperFmtErrorw(s string) ErrflowOption {
	return WrapperFmtErrorf("%s: %w", s, OriginalErr)
}

type originalErrType struct{}

func (err *originalErrType) Error() string {
	return "errflow original error placeholder"
}

// OriginalErr is a placeholder for error in WrapperFmtErrorf(...) call.
var OriginalErr = &originalErrType{}
