package errflow

import (
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatcher_ThenAssignTo(t *testing.T) {
	fn := func(returnErr error) (err error) {
		defer IfError().ThenAssignTo(&err)
		return Check(returnErr)
	}

	assert.Nil(t, fn(nil))
	assert.EqualError(t, fn(fmt.Errorf("test error")), "test error")
}

func TestCatcher_Then(t *testing.T) {
	var outErr error
	fn := func(returnErr error) {
		defer IfError().Then(func(err error) { outErr = err })
		Check(returnErr)
	}

	fn(nil)
	assert.Nil(t, outErr)
	fn(fmt.Errorf("test error"))
	assert.EqualError(t, outErr, "test error")
}

func TestCatcher_ReturnFirst(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnFirst().ThenAssignTo(&err)
		defer Check(fmt.Errorf("second"))
		defer Check(fmt.Errorf("first"))
		return nil
	}

	assert.EqualError(t, fn(), "first")
}

func TestCatcher_ReturnLast(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnLast().ThenAssignTo(&err)
		defer Check(fmt.Errorf("second"))
		defer Check(fmt.Errorf("first"))
		return nil
	}

	assert.EqualError(t, fn(), "second")
}

func TestCatcher_ReturnAll(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnAll().ThenAssignTo(&err)
		defer Check(fmt.Errorf("second"))
		defer Check(fmt.Errorf("first"))
		return nil
	}

	err := fn()
	assert.Contains(t, err.Error(), "first")
	assert.Contains(t, err.Error(), "second")
}

func TestCatcher_LogNone(t *testing.T) {
	var logs []string
	defer SetLogFn(func(s string) {
		logs = append(logs, s)
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogNone().ThenAssignTo(&err)
		defer Check(fmt.Errorf("second"))
		defer Check(fmt.Errorf("first"))
		return nil
	}

	fn()
	assert.Empty(t, logs)
}

func TestCatcher_LogAll(t *testing.T) {
	var logs []string
	defer SetLogFn(func(s string) {
		logs = append(logs, s)
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogAll().ThenAssignTo(&err)
		defer Check(fmt.Errorf("second"))
		defer Check(fmt.Errorf("first"))
		return nil
	}

	fn()
	assert.Equal(t, []string{"first", "second"}, logs)
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

func TestCatcher_UnrelatedPanicStacktrace(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		unrelatedPanicFn()
		return nil
	}

	defer func() {
		recover()
		assert.Contains(t, string(debug.Stack()), "unrelatedPanicFn")
	}()

	fn()
}

func TestCatcher_UnrelatedPanicMultipleErrors(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnFirst().ThenAssignTo(&err)
		defer Check(fmt.Errorf("second"))
		defer Check(fmt.Errorf("first"))
		panic(fmt.Errorf("unrelated panic"))
	}

	assert.PanicsWithError(t, "unrelated panic", func() {
		fn()
	})
}
