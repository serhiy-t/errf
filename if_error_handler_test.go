package errf

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatcher_ThenAssignTo(t *testing.T) {
	fn := func(returnErr error) (err error) {
		defer IfError().ThenAssignTo(&err)
		return TryErr(returnErr)
	}

	assert.Nil(t, fn(nil))
	assert.EqualError(t, fn(fmt.Errorf("test error")), "test error")
}

func TestCatcher_Then(t *testing.T) {
	var outErr error
	fn := func(returnErr error) {
		defer IfError().Then(func(err error) { outErr = err })
		TryErr(returnErr)
	}

	fn(nil)
	assert.Nil(t, outErr)
	fn(fmt.Errorf("test error"))
	assert.EqualError(t, outErr, "test error")
}

func TestCatcher_ThenIgnore(t *testing.T) {
	fn := func(returnErr error) (err error) {
		defer IfError().ThenIgnore()
		return TryErr(returnErr)
	}

	assert.Nil(t, fn(nil))
	assert.Nil(t, fn(fmt.Errorf("test error")))
}

func TestCatcher_ReturnFirst(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnFirst().ThenAssignTo(&err)
		defer TryErr(fmt.Errorf("second"))
		defer TryErr(fmt.Errorf("first"))
		return nil
	}

	assert.EqualError(t, fn(), "first")
}

func TestCatcher_ReturnLast(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnLast().ThenAssignTo(&err)
		defer TryErr(fmt.Errorf("second"))
		defer TryErr(fmt.Errorf("first"))
		return nil
	}

	assert.EqualError(t, fn(), "second")
}

func TestCatcher_ReturnWrapped(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)
		defer TryErr(fmt.Errorf("second"))
		defer TryErr(fmt.Errorf("first"))
		return nil
	}

	err := fn()
	assert.EqualError(t, err, "first (also: second)")
	assert.EqualError(t, errors.Unwrap(err), "first")
}

func TestCatcher_ReturnCombined(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnCombined().ThenAssignTo(&err)
		defer TryErr(fmt.Errorf("second"))
		defer TryErr(fmt.Errorf("first"))
		return nil
	}

	err := fn()
	assert.EqualError(t, err, "combined error {first; second}")
	errs := GetCombinedErrors(err)
	assert.Equal(t, 2, len(errs))
	assert.EqualError(t, errs[0], "first")
	assert.EqualError(t, errs[1], "second")
}

func TestCatcher_LogNever(t *testing.T) {
	var logs []string
	defer SetLogFn(func(m *LogMessage) {
		logs = append(logs, fmt.Sprintf(m.Format, m.A...))
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogNever().ThenAssignTo(&err)
		defer TryErr(fmt.Errorf("second"))
		defer TryErr(fmt.Errorf("first"))
		return nil
	}

	fn()
	assert.Empty(t, logs)
}

func TestCatcher_LogAlways(t *testing.T) {
	var logs []string
	defer SetLogFn(func(m *LogMessage) {
		logs = append(logs, fmt.Sprintf(m.Format, m.A...))
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogAlways().ThenAssignTo(&err)
		defer TryErr(fmt.Errorf("second"))
		defer TryErr(fmt.Errorf("first"))
		return nil
	}

	fn()
	assert.Equal(t, []string{"first", "second"}, logs)
}

func TestCatcher_LogIfSuppressed(t *testing.T) {
	var logs []string
	defer SetLogFn(func(m *LogMessage) {
		logs = append(logs, fmt.Sprintf(m.Format, m.A...))
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogIfSuppressed().ThenAssignTo(&err)
		defer TryErr(fmt.Errorf("second"))
		defer TryErr(fmt.Errorf("first"))
		return nil
	}

	fn()
	assert.Equal(t, []string{"second"}, logs)
}

func unrelatedPanicFn() {
	panic(fmt.Errorf("unrelated panic"))
}

func TestCatcher_UnrelatedPanic(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		unrelatedPanicFn()
		return nil
	}

	assert.PanicsWithError(t, "unrelated panic", func() {
		fn()
	})
}
