package errflow

import "testing"

func Test_getErrorStackTrace(t *testing.T) {
	t.Log("!!!\n" + getErrorStackTrace())
	t.Fail()
}
