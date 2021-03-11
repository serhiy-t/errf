package errflow

import (
	"fmt"
	"log"
	"strings"
)

type LogFn func(logMessage *LogMessage)

var globalLogFn = func(logMessage *LogMessage) {
	var buffer strings.Builder

	for _, tag := range logMessage.Tags {
		fmt.Fprintf(&buffer, "[%s]", tag)
	}
	if buffer.Len() > 0 {
		fmt.Fprintf(&buffer, " ")
	}

	fmt.Fprintf(&buffer, logMessage.Format, logMessage.A...)
	if len(logMessage.Stack) > 0 {
		fmt.Fprintf(&buffer, "\n\nStack:\n%s", logMessage.Stack)
	}

	log.Println(buffer.String())
}

type logFnRestorer struct {
	oldLogFn LogFn
}

func (rlf *logFnRestorer) ThenRestore() {
	globalLogFn = rlf.oldLogFn
}

type LogMessage struct {
	Format string
	A      []interface{}

	Stack string
	Tags  []string
}

// SetLogFn replaces logging function for errflow.
// It returns errflow.DeferRestorer instance,
// which can be used to restore previous logFn, if needed.
// Default log function is log.Println().
func SetLogFn(logFn LogFn) DeferRestorer {
	oldLogFn := globalLogFn
	globalLogFn = logFn
	return &logFnRestorer{
		oldLogFn: oldLogFn,
	}
}
