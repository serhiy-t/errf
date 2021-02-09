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
		Bufio.With(WrapperFmtErrorw("wrapped")).TryWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error")
}

func Test_Bufio_TryBufioWriter(t *testing.T) {
	value := bufio.NewWriter(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Bufio.TryWriter(value, nil))
		Bufio.TryWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_TryBufioReader(t *testing.T) {
	value := bufio.NewReader(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Bufio.TryReader(value, nil))
		Bufio.TryReader(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_TryBufioReadWriter(t *testing.T) {
	value := bufio.NewReadWriter(nil, nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Bufio.TryReadWriter(value, nil))
		Bufio.TryReadWriter(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_TryBufioWriterErr(t *testing.T) {
	value := bufio.NewWriter(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Bufio.TryWriterErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, tryErr)
		Bufio.TryWriterErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_TryBufioReaderErr(t *testing.T) {
	value := bufio.NewReader(nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Bufio.TryReaderErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, tryErr)
		Bufio.TryReaderErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Bufio_TryBufioReadWriterErr(t *testing.T) {
	value := bufio.NewReadWriter(nil, nil)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Bufio.TryReadWriterErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, tryErr)
		Bufio.TryReadWriterErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
