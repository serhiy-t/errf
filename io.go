package errf

import "io"

// Io contains collection of Check* functions for io.* types.
var Io = IoErrflow{}

// IoErrflow implements Check* functions for io package types.
// Clients should not instantiate IoErrflow, use 'errf.Io' instead.
type IoErrflow struct {
	errflow *Errflow
}

// With implements Errflow.With(...) for io types.
func (ef IoErrflow) With(options ...ErrflowOption) IoErrflow {
	return IoErrflow{errflow: ef.errflow.With(options...)}
}

// CheckWriter calls errf.Check and returns a typed value from a function call.
func (ef IoErrflow) CheckWriter(value io.Writer, err error) io.Writer {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckWriteCloser calls errf.Check and returns a typed value from a function call.
func (ef IoErrflow) CheckWriteCloser(value io.WriteCloser, err error) io.WriteCloser {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckReader calls errf.Check and returns a typed value from a function call.
func (ef IoErrflow) CheckReader(value io.Reader, err error) io.Reader {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckReadCloser calls errf.Check and returns a typed value from a function call.
func (ef IoErrflow) CheckReadCloser(value io.ReadCloser, err error) io.ReadCloser {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckWriterErr calls errf.Check and returns a typed value and error from a function call.
func (ef IoErrflow) CheckWriterErr(value io.Writer, err error) (io.Writer, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckWriteCloserErr calls errf.Check and returns a typed value and error from a function call.
func (ef IoErrflow) CheckWriteCloserErr(value io.WriteCloser, err error) (io.WriteCloser, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckReaderErr calls errf.Check and returns a typed value and error from a function call.
func (ef IoErrflow) CheckReaderErr(value io.Reader, err error) (io.Reader, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckReadCloserErr calls errf.Check and returns a typed value and error from a function call.
func (ef IoErrflow) CheckReadCloserErr(value io.ReadCloser, err error) (io.ReadCloser, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}
