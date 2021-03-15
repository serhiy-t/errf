package errf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getErrorStackTrace(t *testing.T) {
	actual := getErrorStackTrace()
	t.Log("Formatted ParsedStack: \n", actual.String())
	assert.Equal(t, 3, len(actual.items))
	assert.Contains(t, actual.items[0].fn, "errflow.Test_getErrorStackTrace")
}
