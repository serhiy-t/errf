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
		Io.With(WrapperFmtErrorw("wrapped")).CheckWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error")
}

func Test_Io_CheckWriter(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Io.CheckWriter(value, nil))
		Io.CheckWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_CheckWriteCloser(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Io.CheckWriteCloser(value, nil))
		Io.CheckWriteCloser(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_CheckReader(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Io.CheckReader(value, nil))
		Io.CheckReader(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_CheckReadCloser(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Io.CheckReadCloser(value, nil))
		Io.CheckReadCloser(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_CheckWriterErr(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Io.CheckWriterErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, checkErr)
		Io.CheckWriterErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_CheckWriteCloserErr(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Io.CheckWriteCloserErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, checkErr)
		Io.CheckWriteCloserErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_CheckReaderErr(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Io.CheckReaderErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, checkErr)
		Io.CheckReaderErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Io_CheckReadCloserErr(t *testing.T) {
	value := &IoWriteReadCloser{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Io.CheckReadCloserErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, checkErr)
		Io.CheckReadCloserErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
