package errf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type IoWriteReadCloser struct{}

func (IoWriteReadCloser) Write(data []byte) (n int, err error) {
	return 0, nil
}
func (IoWriteReadCloser) Read(p []byte) (n int, err error) {
	return 0, nil
}
func (IoWriteReadCloser) Close() error {
	return nil
}

func Test_Io_With(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		Io.With(WrapperFmtErrorw("wrapped")).TryWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error")
}

func Test_Io_TryWriter(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Io.TryWriter(value, nil))
		Io.TryWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_TryWriteCloser(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Io.TryWriteCloser(value, nil))
		Io.TryWriteCloser(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_TryReader(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Io.TryReader(value, nil))
		Io.TryReader(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_TryReadCloser(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Io.TryReadCloser(value, nil))
		Io.TryReadCloser(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_TryWriterErr(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Io.TryWriterErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, tryErr)
		Io.TryWriterErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_TryWriteCloserErr(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Io.TryWriteCloserErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, tryErr)
		Io.TryWriteCloserErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_TryReaderErr(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Io.TryReaderErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, tryErr)
		Io.TryReaderErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_TryReadCloserErr(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Io.TryReadCloserErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, tryErr)
		Io.TryReadCloserErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
