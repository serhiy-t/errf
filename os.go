package errf

import (
	"os"
)

// Os contains collection of Check* functions for os.* types.
var Os = OsErrflow{}

// OsErrflow implements Check* functions for os package types.
//
// Clients should not instantiate OsErrflow, use 'errf.Os' instead.
type OsErrflow struct {
	errflow *Errflow
}

// With implements Errflow.With(...) for os types.
func (ef OsErrflow) With(options ...ErrflowOption) OsErrflow {
	return OsErrflow{errflow: ef.errflow.With(options...)}
}

// CheckFile calls errflow.Check and returns a typed value from a function call.
func (ef OsErrflow) CheckFile(value *os.File, err error) *os.File {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckFileErr calls errflow.Check and returns a typed value and error from a function call.
func (ef OsErrflow) CheckFileErr(value *os.File, err error) (*os.File, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}
