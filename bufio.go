package errf

import (
	"bufio"
)

// Bufio contains collection of Check* functions for bufio.* types.
var Bufio = BufioErrflow{}

// BufioErrflow implements Check* functions for bufio package types.
//
// Clients should not instantiate BufioErrflow, use 'errf.Bufio' instead.
type BufioErrflow struct {
	errflow *Errflow
}

// With implements Errflow.With(...) for bufio types.
func (ef BufioErrflow) With(options ...ErrflowOption) BufioErrflow {
	return BufioErrflow{errflow: ef.errflow.With(options...)}
}

// CheckWriter calls errflow.Check and returns a typed value from a function call.
func (ef BufioErrflow) CheckWriter(value *bufio.Writer, err error) *bufio.Writer {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckReader calls errflow.Check and returns a typed value from a function call.
func (ef BufioErrflow) CheckReader(value *bufio.Reader, err error) *bufio.Reader {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckReadWriter calls errflow.Check and returns a typed value from a function call.
func (ef BufioErrflow) CheckReadWriter(value *bufio.ReadWriter, err error) *bufio.ReadWriter {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckWriterErr calls errflow.Check and returns a typed value and error from a function call.
func (ef BufioErrflow) CheckWriterErr(value *bufio.Writer, err error) (*bufio.Writer, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, nil
}

// CheckReaderErr calls errflow.Check and returns a typed value and error from a function call.
func (ef BufioErrflow) CheckReaderErr(value *bufio.Reader, err error) (*bufio.Reader, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, nil
}

// CheckReadWriterErr calls errflow.Check and returns a typed value and error from a function call.
func (ef BufioErrflow) CheckReadWriterErr(value *bufio.ReadWriter, err error) (*bufio.ReadWriter, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, nil
}
