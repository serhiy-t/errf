package errflow

import "io"

type IoErrflow struct {
	errflow *Errflow
}

func (ef IoErrflow) With(options ...ErrflowOption) IoErrflow {
	return IoErrflow{errflow: ef.errflow.With(options...)}
}

// TryIoWriter calls errflow.Try and returns a typed value from a function call.
func (ef IoErrflow) TryWriter(value io.Writer, err error) io.Writer {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryIoWriteCloser calls errflow.Try and returns a typed value from a function call.
func (ef IoErrflow) TryWriteCloser(value io.WriteCloser, err error) io.WriteCloser {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryIoReader calls errflow.Try and returns a typed value from a function call.
func (ef IoErrflow) TryReader(value io.Reader, err error) io.Reader {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryIoReadCloser calls errflow.Try and returns a typed value from a function call.
func (ef IoErrflow) TryReadCloser(value io.ReadCloser, err error) io.ReadCloser {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// Err

// TryIoWriter calls errflow.Try and returns a typed value from a function call.
func (ef IoErrflow) TryWriterErr(value io.Writer, err error) (io.Writer, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryIoWriteCloser calls errflow.Try and returns a typed value from a function call.
func (ef IoErrflow) TryWriteCloserErr(value io.WriteCloser, err error) (io.WriteCloser, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryIoReader calls errflow.Try and returns a typed value from a function call.
func (ef IoErrflow) TryReaderErr(value io.Reader, err error) (io.Reader, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryIoReadCloser calls errflow.Try and returns a typed value from a function call.
func (ef IoErrflow) TryReadCloserErr(value io.ReadCloser, err error) (io.ReadCloser, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}
