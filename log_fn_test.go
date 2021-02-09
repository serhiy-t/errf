package errf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_defaulGlobalLogFn(t *testing.T) {
	var logs []string
	var fn = defaulGlobalLogFn(func(s string) {
		logs = append(logs, s)
	})

	fn(&LogMessage{
		Format: "hello %s world %d",
		A:      []interface{}{"str", 123},
		Tags:   []string{"my", "tags"},
		Stack:  func() string { return "STACKTRACE" },
	})

	assert.Equal(t, []string{"[my][tags] hello str world 123\n\nStack:\nSTACKTRACE"}, logs)
}

func Test_defaulGlobalLogFn_no_panic(t *testing.T) {
	assert.NotPanics(t, func() {
		globalLogFn(&LogMessage{
			Format: "hello %s world %d",
			A:      []interface{}{"str", 123},
			Tags:   []string{"my", "tags"},
			Stack:  func() string { return "STACKTRACE" },
		})
	})
}
