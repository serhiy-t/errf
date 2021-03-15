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
		Bufio.With(WrapperFmtErrorw("wrapped")).TryBufioWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error")
}

func Test_Bufio_TryBufioWriter(t *testing.T) {
	value := bufio.NewWriter(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Bufio.TryBufioWriter(value, nil))
		Bufio.TryBufioWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_TryBufioReader(t *testing.T) {
	value := bufio.NewReader(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Bufio.TryBufioReader(value, nil))
		Bufio.TryBufioReader(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_TryBufioReadWriter(t *testing.T) {
	value := bufio.NewReadWriter(nil, nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Bufio.TryBufioReadWriter(value, nil))
		Bufio.TryBufioReadWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_TryBufioWriterErr(t *testing.T) {
	value := bufio.NewWriter(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Bufio.TryBufioWriterErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, tryErr)
		Bufio.TryBufioWriterErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_TryBufioReaderErr(t *testing.T) {
	value := bufio.NewReader(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Bufio.TryBufioReaderErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, tryErr)
		Bufio.TryBufioReaderErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_TryBufioReadWriterErr(t *testing.T) {
	value := bufio.NewReadWriter(nil, nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Bufio.TryBufioReadWriterErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, tryErr)
		Bufio.TryBufioReadWriterErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
