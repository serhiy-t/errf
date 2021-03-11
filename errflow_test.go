package errflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrflow_TryErr(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		TryErr(fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_TryAny(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, "value", TryAny("value", fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_TryDiscard(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		TryDiscard("value", fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestErrflow_Log(t *testing.T) {
	var logs []string
	defer SetLogFn(func(logMessage *LogMessage) {
		logs = append(logs, fmt.Sprintf(logMessage.Format, logMessage.A...))
	}).ThenRestore()

	fn := func() (err error) {
		defer IfError().LogNever().ThenAssignTo(&err)
		defer Log(fmt.Errorf("error message"))
		return nil
	}

	assert.NoError(t, fn())
	assert.Equal(t, []string{"error message"}, logs)
}

func TestErrflow_ErrorIf(t *testing.T) {
	fn := func(returnErr bool) (err error) {
		defer IfError().ThenAssignTo(&err)

		return TryCondition(returnErr, "test error")
	}

	assert.Nil(t, fn(false))
	assert.EqualError(t, fn(true), "test error")
}

func TestErrflow_Error_With(t *testing.T) {
	var logs []string
	defer SetLogFn(func(logMessage *LogMessage) {
		logs = append(logs, fmt.Sprintf(logMessage.Format, logMessage.A...))
	}).ThenRestore()

	fn := func(options ...ErrflowOption) (err error) {
		defer IfError().ThenAssignTo(&err)

		defer With(options...).TryErr(fmt.Errorf("error3"))
		defer With(options...).TryErr(fmt.Errorf("error2"))

		return With(options...).TryErr(fmt.Errorf("error1"))
	}

	assert.EqualError(t, fn(ReturnStrategyFirst, LogStrategyAlways), "error1")
	assert.Equal(t, []string{"error1", "error2", "error3"}, logs)
	logs = nil

	assert.EqualError(t, fn(ReturnStrategyLast, LogStrategyIfSuppressed), "error3")
	assert.Equal(t, []string{"error1", "error2"}, logs)
	logs = nil

	assert.EqualError(t, fn(ReturnStrategyWrapped),
		"error1 (also: error2) (also: error3)")
	assert.Empty(t, logs)

	assert.EqualError(t, fn(ReturnStrategyCombined),
		"combined error {error1; error2; error3}")
}

func TestErrflow_IfError_Apply(t *testing.T) {
	var logs []string
	defer SetLogFn(func(logMessage *LogMessage) {
		logs = append(logs, fmt.Sprintf(logMessage.Format, logMessage.A...))
	}).ThenRestore()

	fn := func(options ...ErrflowOption) (err error) {
		defer IfError().Apply(options...).ThenAssignTo(&err)

		defer TryErr(fmt.Errorf("error3"))
		defer TryErr(fmt.Errorf("error2"))

		return TryErr(fmt.Errorf("error1"))
	}

	assert.EqualError(t, fn(ReturnStrategyFirst, LogStrategyAlways), "error1")
	assert.Equal(t, []string{"error1", "error2", "error3"}, logs)
	logs = nil

	assert.EqualError(t, fn(ReturnStrategyLast, LogStrategyIfSuppressed), "error3")
	assert.Equal(t, []string{"error1", "error2"}, logs)
	logs = nil

	assert.EqualError(t, fn(ReturnStrategyWrapped),
		"error1 (also: error2) (also: error3)")
	assert.Empty(t, logs)

	assert.EqualError(t, fn(ReturnStrategyCombined),
		"combined error {error1; error2; error3}")
}
