package errf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_CorrectCheckErr(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		return CheckErr(nil)
	}
	assert.NotPanics(t, func() {
		assert.NoError(t, fn())
	})
}

func TestValidator_CheckErrWithoutIfError(t *testing.T) {
	fn := func() (err error) {
		return CheckErr(nil)
	}
	assert.PanicsWithError(t, "errflow incorrect call sequence", func() {
		fn()
	})
}

func TestValidator_CheckErrWithoutErrorDisabledValidator(t *testing.T) {
	defer SetNoopValidator().ThenRestore()

	fn := func() (err error) {
		return CheckErr(nil)
	}
	assert.NotPanics(t, func() {
		fn()
	})
}

func TestValidator_DisabledValidator(t *testing.T) {
	defer SetNoopValidator().ThenRestore()

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		return CheckErr(nil)
	}
	assert.NoError(t, fn())
	assert.NotPanics(t, func() {
		fn()
	})
}

func TestValidator_NestedCatches(t *testing.T) {
	var fn func(level int) error
	fn = func(level int) (err error) {
		defer IfError().ThenAssignTo(&err)

		if level > 0 {
			return fn(level - 1)
		}
		return CheckErr(fmt.Errorf("error message"))
	}
	assert.EqualError(t, fn(5), "error message")
}

func TestValidator_IncorrectNestedFns(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)

		func() {
			CheckErr(fmt.Errorf("error message"))
		}()

		return nil
	}
	assert.PanicsWithError(t, "errflow incorrect call sequence", func() {
		fn()
	})
}

func TestValidator_CorrectNestedFns(t *testing.T) {
	fn := func() (err error) {
		return func() (err error) {
			defer IfError().ThenAssignTo(&err)

			return CheckErr(fmt.Errorf("error message"))
		}()
	}
	assert.EqualError(t, fn(), "error message")
}

func TestValidator_MissingCatchStatement(t *testing.T) {
	var fn func(level int) error
	fn = func(level int) (err error) {
		defer IfError()

		if level > 0 {
			return fn(level - 1)
		}
		return CheckErr(fmt.Errorf("error message"))
	}
	assert.PanicsWithError(t, "errflow incorrect call sequence", func() {
		fn(5)
	})
}
