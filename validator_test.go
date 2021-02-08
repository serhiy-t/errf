package errflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_CheckWithoutCatch(t *testing.T) {
	fn := func() (err error) {
		return C(nil)
	}
	assert.PanicsWithError(t, "errflow incorrect call sequence", func() {
		fn()
	})
}

func TestValidator_CheckWithoutCatchDisabledValidator(t *testing.T) {
	defer SetNoopValidator().Close()

	fn := func() (err error) {
		return C(nil)
	}
	assert.NotPanics(t, func() {
		fn()
	})
}

func TestValidator_DisabledValidator(t *testing.T) {
	defer SetNoopValidator().Close()

	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		return C(nil)
	}
	assert.NoError(t, fn())
	assert.NotPanics(t, func() {
		fn()
	})
}

func TestValidator_NestedCatches(t *testing.T) {
	var fn func(level int) error
	fn = func(level int) (err error) {
		defer Catch().WriteTo(&err)

		if level > 0 {
			return fn(level - 1)
		}
		return C(fmt.Errorf("error message"))
	}
	assert.EqualError(t, fn(5), "error message")
}

func TestValidator_MissingCatchStatement(t *testing.T) {
	var fn func(level int) error
	fn = func(level int) (err error) {
		defer Catch()

		if level > 0 {
			return fn(level - 1)
		}
		return C(fmt.Errorf("error message"))
	}
	assert.PanicsWithError(t, "errflow incorrect call sequence", func() {
		fn(5)
	})
}
