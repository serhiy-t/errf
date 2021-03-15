package errf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getErrorStackTrace(t *testing.T) {
	actual := getErrorStackTrace()
	t.Log("Formatted ParsedStack: \n", actual.String())
	assert.Equal(t, 3, len(actual.Items))
	assert.Contains(t, actual.Items[0].Fn, "errflow.Test_getErrorStackTrace")
}
