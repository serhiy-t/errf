package errf

import (
	"bufio"
)

type BufioErrflow struct {
	errflow *Errflow
}

func (ef BufioErrflow) With(options ...ErrflowOption) BufioErrflow {
	return BufioErrflow{errflow: ef.errflow.With(options...)}
}

// TryBufioWriter calls errflow.Try and returns a typed value from a function call.
func (ef BufioErrflow) TryBufioWriter(value *bufio.Writer, err error) *bufio.Writer {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryBufioReader calls errflow.Try and returns a typed value from a function call.
func (ef BufioErrflow) TryBufioReader(value *bufio.Reader, err error) *bufio.Reader {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryBufioReadWriter calls errflow.Try and returns a typed value from a function call.
func (ef BufioErrflow) TryBufioReadWriter(value *bufio.ReadWriter, err error) *bufio.ReadWriter {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// Errs

// TryBufioWriter calls errflow.Try and returns a typed value from a function call.
func (ef BufioErrflow) TryBufioWriterErr(value *bufio.Writer, err error) (*bufio.Writer, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, nil
}

// TryBufioReader calls errflow.Try and returns a typed value from a function call.
func (ef BufioErrflow) TryBufioReaderErr(value *bufio.Reader, err error) (*bufio.Reader, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, nil
}

// TryBufioReadWriter calls errflow.Try and returns a typed value from a function call.
func (ef BufioErrflow) TryBufioReadWriterErr(value *bufio.ReadWriter, err error) (*bufio.ReadWriter, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, nil
}
