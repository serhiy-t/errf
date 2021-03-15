package errf

import (
	"bufio"
)

// Bufio contains collection of Try* functions for bufio.* types.
var Bufio = BufioErrflow{}

// BufioErrflow implements Try* functions for bufio package types.
// Clients should not instantiate BufioErrflow, use 'errf.Bufio' instead.
type BufioErrflow struct {
	errflow *Errflow
}

// With implements Errflow.With(...) for bufio types.
func (ef BufioErrflow) With(options ...ErrflowOption) BufioErrflow {
	return BufioErrflow{errflow: ef.errflow.With(options...)}
}

// TryWriter calls errflow.Try and returns a typed value from a function call.
func (ef BufioErrflow) TryWriter(value *bufio.Writer, err error) *bufio.Writer {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryReader calls errflow.Try and returns a typed value from a function call.
func (ef BufioErrflow) TryReader(value *bufio.Reader, err error) *bufio.Reader {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryReadWriter calls errflow.Try and returns a typed value from a function call.
func (ef BufioErrflow) TryReadWriter(value *bufio.ReadWriter, err error) *bufio.ReadWriter {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryWriterErr calls errflow.Try and returns a typed value and error from a function call.
func (ef BufioErrflow) TryWriterErr(value *bufio.Writer, err error) (*bufio.Writer, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, nil
}

// TryReaderErr calls errflow.Try and returns a typed value and error from a function call.
func (ef BufioErrflow) TryReaderErr(value *bufio.Reader, err error) (*bufio.Reader, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, nil
}

// TryReadWriterErr calls errflow.Try and returns a typed value and error from a function call.
func (ef BufioErrflow) TryReadWriterErr(value *bufio.ReadWriter, err error) (*bufio.ReadWriter, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, nil
}
