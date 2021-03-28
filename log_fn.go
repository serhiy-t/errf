package errf

import (
	"fmt"
	"log"
	"strings"
)

// LogFn defines logger function API for errflow.
type LogFn func(logMessage *LogMessage)

var globalLogFn = defaultGlobalLogFn(func(s string) {
	log.Println(s)
})

func defaultGlobalLogFn(printFn func(string)) LogFn {
	return func(logMessage *LogMessage) {
		var buffer strings.Builder

		for _, tag := range logMessage.Tags {
			_, _ = fmt.Fprintf(&buffer, "[%s]", tag)
		}
		if buffer.Len() > 0 {
			_, _ = fmt.Fprintf(&buffer, " ")
		}

		_, _ = fmt.Fprintf(&buffer, logMessage.Format, logMessage.A...)
		if logMessage.Stack != nil {
			_, _ = fmt.Fprintf(&buffer, "\n\nStack:\n%s", logMessage.Stack())
		}

		printFn(buffer.String())
	}
}

type logFnRestorer struct {
	oldLogFn LogFn
}

func (rlf *logFnRestorer) ThenRestore() {
	globalLogFn = rlf.oldLogFn
}

// LogMessage defines a single log message interface.
type LogMessage struct {
	// Format is a format string.
	Format string
	// A contains arguments for Format string.
	A []interface{}

	// Stack function returns stacktrace string.
	Stack func() string

	// Tags contains additional tags.
	Tags []string
}

// SetLogFn replaces logging function for errflow.
// It returns errf.DeferRestorer instance,
// which can be used to restore previous logFn, if needed.
// Default log function is log.Println().
func SetLogFn(logFn LogFn) DeferRestorer {
	oldLogFn := globalLogFn
	globalLogFn = logFn
	return &logFnRestorer{
		oldLogFn: oldLogFn,
	}
}
