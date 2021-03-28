package errf

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Os_With(t *testing.T) {
	value := &os.File{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		Os.With(WrapperFmtErrorw("wrapped")).CheckFile(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error")
}

func Test_Os_CheckFile(t *testing.T) {
	value := &os.File{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Same(t, value, Os.CheckFile(value, nil))
		Os.CheckFile(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_Os_CheckFileErr(t *testing.T) {
	value := &os.File{}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Os.CheckFileErr(value, nil)
		assert.Same(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Os.CheckFileErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
