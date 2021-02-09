package errflow

import (
	"log"
)

var globalLogFn = func(s string) { log.Print(s) }

type logFnRestorer struct {
	oldLogFn func(s string)
}

func (rlf *logFnRestorer) ThenRestore() {
	globalLogFn = rlf.oldLogFn
}

// SetLogFn replaces logging function for errflow.
// It returns errflow.DeferRestorer instance,
// which can be used to restore previous logFn, if needed.
// Default log function is log.Println().
func SetLogFn(logFn func(s string)) DeferRestorer {
	oldLogFn := globalLogFn
	globalLogFn = logFn
	return &logFnRestorer{
		oldLogFn: oldLogFn,
	}
}
