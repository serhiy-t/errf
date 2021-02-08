package errflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatcher_WriteTo(t *testing.T) {
	fn := func(returnErr error) (err error) {
		defer Catch().WriteTo(&err)
		return C(returnErr)
	}

	assert.Nil(t, fn(nil))
	assert.EqualError(t, fn(fmt.Errorf("test error")), "test error")
}

func TestCatcher_Then(t *testing.T) {
	var outErr error
	fn := func(returnErr error) {
		defer Catch().Then(func(err error) { outErr = err })
		C(returnErr)
	}

	fn(nil)
	assert.Nil(t, outErr)
	fn(fmt.Errorf("test error"))
	assert.EqualError(t, outErr, "test error")
}

func TestCatcher_ReturnFirst(t *testing.T) {
	fn := func() (err error) {
		defer Catch().ReturnFirst().WriteTo(&err)
		defer C(fmt.Errorf("second"))
		defer C(fmt.Errorf("first"))
		return nil
	}

	assert.EqualError(t, fn(), "first")
}

func TestCatcher_ReturnLast(t *testing.T) {
	fn := func() (err error) {
		defer Catch().ReturnLast().WriteTo(&err)
		defer C(fmt.Errorf("second"))
		defer C(fmt.Errorf("first"))
		return nil
	}

	assert.EqualError(t, fn(), "second")
}

func TestCatcher_ReturnAll(t *testing.T) {
	fn := func() (err error) {
		defer Catch().ReturnAll().WriteTo(&err)
		defer C(fmt.Errorf("second"))
		defer C(fmt.Errorf("first"))
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
	}).Close()

	fn := func() (err error) {
		defer Catch().LogNone().WriteTo(&err)
		defer C(fmt.Errorf("second"))
		defer C(fmt.Errorf("first"))
		return nil
	}

	fn()
	assert.Empty(t, logs)
}

func TestCatcher_LogAll(t *testing.T) {
	var logs []string
	defer SetLogFn(func(s string) {
		logs = append(logs, s)
	}).Close()

	fn := func() (err error) {
		defer Catch().LogAll().WriteTo(&err)
		defer C(fmt.Errorf("second"))
		defer C(fmt.Errorf("first"))
		return nil
	}

	fn()
	assert.Equal(t, []string{"first", "second"}, logs)
}

func TestCatcher_UnrelatedPanic(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		panic(fmt.Errorf("unrelated panic"))
	}

	assert.PanicsWithError(t, "unrelated panic", func() {
		fn()
	})
}

func TestCatcher_UnrelatedPanicMultipleErrors(t *testing.T) {
	fn := func() (err error) {
		defer Catch().ReturnFirst().WriteTo(&err)
		defer C(fmt.Errorf("second"))
		defer C(fmt.Errorf("first"))
		panic(fmt.Errorf("unrelated panic"))
	}

	assert.PanicsWithError(t, "unrelated panic", func() {
		fn()
	})
}
