package errflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IfErrorAssignTo_Default(t *testing.T) {
	fn := func() (err error) {
		defer IfErrorAssignTo(&err, fmt.Errorf("error2"))
		defer IfErrorAssignTo(&err, fmt.Errorf("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "error1")
}

func Test_IfErrorAssignTo_ReturnStrategyLast_panics(t *testing.T) {
	fn := func() (err error) {
		defer With(ReturnStrategyLast).IfErrorAssignTo(&err, nil)
		return nil
	}

	assert.Panics(t, func() {
		fn()
	})
}

func Test_IfErrorAssignTo_ReturnStrategyWrapped(t *testing.T) {
	fn := func() (err error) {
		defer With(ReturnStrategyWrapped).IfErrorAssignTo(&err, fmt.Errorf("error3"))
		defer With(ReturnStrategyWrapped).IfErrorAssignTo(&err, fmt.Errorf("error2"))
		defer IfErrorAssignTo(&err, fmt.Errorf("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "error1 (also: error2) (also: error3)")
}

func Test_IfErrorAssignTo_ReturnStrategy(t *testing.T) {
	fn := func() (err error) {
		defer With(ReturnStrategyWrapped).IfErrorAssignTo(&err, fmt.Errorf("error3"))
		defer With(ReturnStrategyWrapped).IfErrorAssignTo(&err, fmt.Errorf("error2"))
		defer IfErrorAssignTo(&err, fmt.Errorf("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "error1 (also: error2) (also: error3)")
}

func Test_IfErrorAssignTo_LogAll(t *testing.T) {
	var logs []string
	defer SetLogFn(func(logMessage *LogMessage) {
		logs = append(logs, fmt.Sprintf(logMessage.Format, logMessage.A...))
	}).ThenRestore()

	fn := func() (err error) {
		defer IfErrorAssignTo(&err, fmt.Errorf("error3"))
		defer With(LogStrategyAlways).IfErrorAssignTo(&err, fmt.Errorf("error2"))
		defer With(LogStrategyAlways).IfErrorAssignTo(&err, fmt.Errorf("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "error1")
	assert.Equal(t, []string{"error1", "error2"}, logs)
}

func Test_IfErrorAssignTo_LogIfSuppressed(t *testing.T) {
	var logs []string
	defer SetLogFn(func(logMessage *LogMessage) {
		logs = append(logs, fmt.Sprintf(logMessage.Format, logMessage.A...))
	}).ThenRestore()

	fn := func() (err error) {
		defer With(
			LogStrategyIfSuppressed, ReturnStrategyFirst,
		).IfErrorAssignTo(&err, fmt.Errorf("error2"))
		defer IfErrorAssignTo(&err, fmt.Errorf("error1"))
		return nil
	}

	assert.EqualError(t, fn(), "error1")
	assert.Equal(t, []string{"error2"}, logs)
}
