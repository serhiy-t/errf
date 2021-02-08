package errflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrflow_Check(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		Check(fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_C(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		C(fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_Untyped(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, "value", Untyped("value", fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_U(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, "value", U("value", fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_IgnoreReturnValue(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		IgnoreReturnValue("value", fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_I(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		I("value", fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
