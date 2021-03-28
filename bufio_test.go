package errf

import (
	"bufio"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bufio_With(t *testing.T) {
	value := bufio.NewWriter(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		Bufio.With(WrapperFmtErrorw("wrapped")).CheckWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error")
}

func Test_Bufio_CheckBufioWriter(t *testing.T) {
	value := bufio.NewWriter(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Bufio.CheckWriter(value, nil))
		Bufio.CheckWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_CheckBufioReader(t *testing.T) {
	value := bufio.NewReader(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Bufio.CheckReader(value, nil))
		Bufio.CheckReader(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_CheckBufioReadWriter(t *testing.T) {
	value := bufio.NewReadWriter(nil, nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Bufio.CheckReadWriter(value, nil))
		Bufio.CheckReadWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_CheckBufioWriterErr(t *testing.T) {
	value := bufio.NewWriter(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Bufio.CheckWriterErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Bufio.CheckWriterErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_CheckBufioReaderErr(t *testing.T) {
	value := bufio.NewReader(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Bufio.CheckReaderErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Bufio.CheckReaderErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_CheckBufioReadWriterErr(t *testing.T) {
	value := bufio.NewReadWriter(nil, nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Bufio.CheckReadWriterErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Bufio.CheckReadWriterErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
