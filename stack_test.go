package errf

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getErrorStackTrace(t *testing.T) {
	actual := getErrorStackTrace()
	t.Log("Formatted ParsedStack: \n", actual.String())
	assert.Equal(t, 3, len(actual.items))
	assert.Contains(t, actual.items[0].fn, "errf.Test_getErrorStackTrace")
}

func Test_getErrorStackTrace_panicErrors(t *testing.T) {
	var stacks []string
	defer SetLogFn(func(logMessage *LogMessage) {
		stacks = append(stacks, logMessage.Stack())
	}).ThenRestore()

	fn := func() error {
		defer IfError().LogAlways().ThenIgnore()
		defer CheckErr(
			fmt.Errorf("error 3"))
		defer CheckErr(
			fmt.Errorf("error 2"))
		CheckErr(fmt.Errorf("error 1"))
		return nil
	}

	_ = fn()

	assert.Equal(t, 3, len(stacks))
	assert.Contains(t, strings.Split(strings.TrimSpace(stacks[0]), "\n")[2], "errf/stack_test.go:30")
	assert.Contains(t, strings.Split(strings.TrimSpace(stacks[1]), "\n")[2], "errf/stack_test.go:30")
	assert.Contains(t, strings.Split(strings.TrimSpace(stacks[2]), "\n")[2], "errf/stack_test.go:30")
}
