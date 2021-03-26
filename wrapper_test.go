package errf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Wrapper(t *testing.T) {
	fn := func() (err error) {
		defer IfError().Apply(Wrapper(func(err error) error {
			return fmt.Errorf("wrapped: %w", err)
		})).ThenAssignTo(&err)

		CheckErr(fmt.Errorf("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error1")
}

func Test_NestedWrappers(t *testing.T) {
	fn := func() (err error) {
		defer IfError().Apply(Wrapper(func(err error) error {
			return fmt.Errorf("wrapped2: %w", err)
		})).ThenAssignTo(&err)

		With(Wrapper(func(err error) error {
			return fmt.Errorf("wrapped1: %w", err)
		})).CheckErr(fmt.Errorf("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped2: wrapped1: error1")
}

func Test_Wrapper_IfErrorAssignTo(t *testing.T) {
	fn := func() (err error) {
		defer With(Wrapper(func(err error) error {
			return fmt.Errorf("wrapped: %w", err)
		})).IfErrorAssignTo(&err, errorFn("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error1")
}

func Test_Wrapper_nil(t *testing.T) {
	fn := func() (err error) {
		defer IfError().Apply(Wrapper(nil)).ThenAssignTo(&err)

		CheckErr(fmt.Errorf("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "error1")
}

func Test_WrapperFmtErrorf(t *testing.T) {
	fn := func() (err error) {
		defer IfError().Apply(
			WrapperFmtErrorf("wrapped: %w", OriginalErr),
		).ThenAssignTo(&err)

		CheckErr(fmt.Errorf("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error1")
}

func Test_WrapperFmtErrorW(t *testing.T) {
	fn := func() (err error) {
		defer IfError().Apply(WrapperFmtErrorw("wrapped")).ThenAssignTo(&err)

		CheckErr(fmt.Errorf("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error1")
}

func Test_OriginalErr(t *testing.T) {
	assert.EqualError(t, OriginalErr, "errflow original error placeholder")
}
