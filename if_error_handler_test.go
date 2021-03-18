package errf

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfErrorHandler_ThenAssignTo(t *testing.T) {
	fn := func(returnErr error) (err error) {
		defer IfError().ThenAssignTo(&err)
		return CheckErr(returnErr)
	}

	assert.Nil(t, fn(nil))
	assert.EqualError(t, fn(fmt.Errorf("test error")), "test error")
}

func TestIfErrorHandler_Then(t *testing.T) {
	var outErr error
	fn := func(returnErr error) {
		defer IfError().Then(func(err error) { outErr = err })
		CheckErr(returnErr)
	}

	fn(nil)
	assert.Nil(t, outErr)
	fn(fmt.Errorf("test error"))
	assert.EqualError(t, outErr, "test error")
}

func TestIfErrorHandler_Then_Multiple(t *testing.T) {
	var errs []error
	fn := func(returnErr error) {
		defer IfError().Then(
			func(err error) {
				errs = append(errs, fmt.Errorf("then 1: %w", err))
			},
			func(err error) {
				errs = append(errs, fmt.Errorf("then 2: %w", err))
			},
			func(err error) {
				errs = append(errs, fmt.Errorf("then 3: %w", err))
			},
		)
		CheckErr(returnErr)
	}

	fn(nil)
	assert.Empty(t, errs)
	fn(fmt.Errorf("test error"))
	assert.Len(t, errs, 3)
	assert.EqualError(t, errs[0], "then 1: test error")
	assert.EqualError(t, errs[1], "then 2: test error")
	assert.EqualError(t, errs[2], "then 3: test error")
}

func TestIfErrorHandler_ThenIgnore(t *testing.T) {
	fn := func(returnErr error) (err error) {
		defer IfError().ThenIgnore()
		return CheckErr(returnErr)
	}

	assert.Nil(t, fn(nil))
	assert.Nil(t, fn(fmt.Errorf("test error")))
}

func TestIfErrorHandler_ReturnFirst(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnFirst().ThenAssignTo(&err)
		defer CheckErr(fmt.Errorf("second"))
		defer CheckErr(fmt.Errorf("first"))
		return nil
	}

	assert.EqualError(t, fn(), "first")
}

func TestIfErrorHandler_ReturnLast(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnLast().ThenAssignTo(&err)
		defer CheckErr(fmt.Errorf("second"))
		defer CheckErr(fmt.Errorf("first"))
		return nil
	}

	assert.EqualError(t, fn(), "second")
}

func TestIfErrorHandler_ReturnWrapped(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)
		defer CheckErr(fmt.Errorf("second"))
		defer CheckErr(fmt.Errorf("first"))
		return nil
	}

	err := fn()
	assert.EqualError(t, err, "first (also: second)")
	assert.EqualError(t, errors.Unwrap(err), "first")
}

func TestIfErrorHandler_ReturnCombined(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnCombined().ThenAssignTo(&err)
		defer CheckErr(fmt.Errorf("second"))
		defer CheckErr(fmt.Errorf("first"))
		return nil
	}

	err := fn()
	assert.EqualError(t, err, "combined error {first; second}")
	errs := GetCombinedErrors(err)
	assert.Equal(t, 2, len(errs))
	assert.EqualError(t, errs[0], "first")
	assert.EqualError(t, errs[1], "second")
}

func TestIfErrorHandler_LogNever(t *testing.T) {
	var logs []string
	defer SetLogFn(func(m *LogMessage) {
		logs = append(logs, fmt.Sprintf(m.Format, m.A...))
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogNever().ThenAssignTo(&err)
		defer CheckErr(fmt.Errorf("second"))
		defer CheckErr(fmt.Errorf("first"))
		return nil
	}

	fn()
	assert.Empty(t, logs)
}

func TestIfErrorHandler_LogAlways(t *testing.T) {
	var logs []string
	defer SetLogFn(func(m *LogMessage) {
		logs = append(logs, fmt.Sprintf(m.Format, m.A...))
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogAlways().ThenAssignTo(&err)
		defer CheckErr(fmt.Errorf("second"))
		defer CheckErr(fmt.Errorf("first"))
		return nil
	}

	fn()
	assert.Equal(t, []string{"first", "second"}, logs)
}

func TestIfErrorHandler_LogIfSuppressed(t *testing.T) {
	var logs []string
	defer SetLogFn(func(m *LogMessage) {
		logs = append(logs, fmt.Sprintf(m.Format, m.A...))
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogIfSuppressed().ThenAssignTo(&err)
		defer CheckErr(fmt.Errorf("second"))
		defer CheckErr(fmt.Errorf("first"))
		return nil
	}

	fn()
	assert.Equal(t, []string{"second"}, logs)
}

func unrelatedPanicFn() {
	panic(fmt.Errorf("unrelated panic"))
}

func TestIfErrorHandler_UnrelatedPanic(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		unrelatedPanicFn()
		return nil
	}

	assert.PanicsWithError(t, "unrelated panic", func() {
		fn()
	})
}

func TestIfErrorHandler_ThenAssignTo_ExternalError(t *testing.T) {
	var logs []string
	defer SetLogFn(func(m *LogMessage) {
		logs = append(logs, fmt.Sprintf(m.Format, m.A...))
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogIfSuppressed().ThenAssignTo(&err)
		defer CheckErr(fmt.Errorf("second"))
		return fmt.Errorf("first")
	}

	assert.EqualError(t, fn(), "second")
	assert.Equal(t, []string{"first"}, logs)
}
