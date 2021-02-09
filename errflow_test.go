package errflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrflow_Check(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		Check(fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_CheckUntyped(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, "value", CheckUntyped("value", fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_CheckIgnoreValue(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		CheckIgnoreValue("value", fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_OnlyLog(t *testing.T) {
	var logs []string
	defer SetLogFn(func(s string) {
		logs = append(logs, s)
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogNone().ThenAssignTo(&err)
		defer OnlyLog(fmt.Errorf("error message"))
		return nil
	}

	assert.NoError(t, fn())
	assert.Equal(t, []string{"error message"}, logs)
}
