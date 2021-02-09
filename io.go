package errf

import "io"

// Io contains collection of Try* functions for io.* types.
var Io = IoErrflow{}

// IoErrflow implements Try* functions for io package types.
// Clients should not instantiate IoErrflow, use 'errf.Io' instead.
type IoErrflow struct {
	errflow *Errflow
}

// With implements Errflow.With(...) for io types.
func (ef IoErrflow) With(options ...ErrflowOption) IoErrflow {
	return IoErrflow{errflow: ef.errflow.With(options...)}
}

// TryWriter calls errf.Try and returns a typed value from a function call.
func (ef IoErrflow) TryWriter(value io.Writer, err error) io.Writer {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryWriteCloser calls errf.Try and returns a typed value from a function call.
func (ef IoErrflow) TryWriteCloser(value io.WriteCloser, err error) io.WriteCloser {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryReader calls errf.Try and returns a typed value from a function call.
func (ef IoErrflow) TryReader(value io.Reader, err error) io.Reader {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryReadCloser calls errf.Try and returns a typed value from a function call.
func (ef IoErrflow) TryReadCloser(value io.ReadCloser, err error) io.ReadCloser {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryWriterErr calls errf.Try and returns a typed value and error from a function call.
func (ef IoErrflow) TryWriterErr(value io.Writer, err error) (io.Writer, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryWriteCloserErr calls errf.Try and returns a typed value and error from a function call.
func (ef IoErrflow) TryWriteCloserErr(value io.WriteCloser, err error) (io.WriteCloser, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryReaderErr calls errf.Try and returns a typed value and error from a function call.
func (ef IoErrflow) TryReaderErr(value io.Reader, err error) (io.Reader, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryReadCloserErr calls errf.Try and returns a typed value and error from a function call.
func (ef IoErrflow) TryReadCloserErr(value io.ReadCloser, err error) (io.ReadCloser, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}
