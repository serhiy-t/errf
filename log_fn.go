package errflow

import (
	"io"
	"log"
)

var globalLogFn = func(s string) { log.Print(s) }

type restoreLogFn struct {
	oldLogFn func(s string)
}

func (rlf *restoreLogFn) Close() error {
	globalLogFn = rlf.oldLogFn
	return nil
}

// SetLogFn replaces logging function for errflow.
// It returns io.Closer instance,
// which can be used to restore previous logFn,
// but also can be ignored, if not needed.
// Default log function is log.Println().
func SetLogFn(logFn func(s string)) io.Closer {
	oldLogFn := globalLogFn
	globalLogFn = logFn
	return &restoreLogFn{
		oldLogFn: oldLogFn,
	}
}
