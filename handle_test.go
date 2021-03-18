package errf

import (
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PanicErr_Error(t *testing.T) {
	assert.Equal(t, "panic: hello", PanicErr{PanicObj: "hello"}.Error())
}

type errPtrType struct{}

func (et *errPtrType) Error() string { return "errPtrType" }

type errType struct{}

func (et errType) Error() string { return "errType" }

func testSuccess() {

}

func testPanic() {
	panic("test panic")
}

var errTestErrorInstance = fmt.Errorf("test error")

func testErrorSkipInErrfStackTrace() {
	panic(errflowThrow{
		items: []errflowThrowItem{
			{
				ef:  DefaultErrflow,
				err: errTestErrorInstance,
			},
		},
	})
}

type testHelper struct {
	err       error
	wasCalled bool
}

func (h *testHelper) clear() {
	h.err = nil
	h.wasCalled = false
}

func (h *testHelper) calledErr(err error) {
	h.err = err
	h.wasCalled = true
}

func (h *testHelper) calledPanic(panicObj interface{}) {
	h.err = PanicErr{PanicObj: panicObj}
	h.wasCalled = true
}

func (h *testHelper) called() {
	h.err = nil
	h.wasCalled = true
}

func assertCalled(t *testing.T, h *testHelper) {
	assert.True(t, h.wasCalled)
	assert.Nil(t, h.err)
}

func assertCalledTestErr(t *testing.T, h *testHelper) {
	assert.True(t, h.wasCalled)
	assert.EqualError(t, h.err, errTestErrorInstance.Error())
}

func assertCalledTestPanic(t *testing.T, h *testHelper) {
	assert.True(t, h.wasCalled)
	assert.EqualError(t, h.err, "panic: test panic")
}

func assertNotCalled(t *testing.T, h *testHelper) {
	assert.False(t, h.wasCalled)
}

type assertFn func(t *testing.T, th *testHelper)

type testConfig struct {
	successAssert assertFn
	errorAssert   assertFn
	panicAssert   assertFn
}

func doTest(t *testing.T, config testConfig, testFn func(*testHelper, func())) {
	var th testHelper

	if config.successAssert != nil {
		th.clear()
		testFn(&th, testSuccess)
		config.successAssert(t, &th)
	}

	if config.errorAssert != nil {
		th.clear()
		testFn(&th, testErrorSkipInErrfStackTrace)
		config.errorAssert(t, &th)
	}

	if config.panicAssert != nil {
		th.clear()
		assert.Panics(t, func() { testFn(&th, testPanic) })
		config.panicAssert(t, &th)
	}
}

func Test_Handler_Always(t *testing.T) {
	doTest(t, testConfig{
		successAssert: assertCalled,
		errorAssert:   assertCalled,
		panicAssert:   assertCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().Always(func() {
			th.called()
		})

		producer()
	})
}

func Test_Handler_Everything(t *testing.T) {
	doTest(t, testConfig{
		successAssert: assertCalled,
		errorAssert:   assertCalledTestErr,
		panicAssert:   assertCalledTestPanic,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().Everything(func(err error) {
			th.calledErr(err)
		})

		producer()
	})
}

func Test_Handler_OnErr(t *testing.T) {
	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertCalledTestErr,
		panicAssert:   assertNotCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnErr(func(err error) {
			th.calledErr(err)
		})

		producer()
	})
}

func Test_Handler_OnErrIs(t *testing.T) {
	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertNotCalled,
		panicAssert:   assertNotCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnErrIs(io.EOF, func() {
			th.called()
		})

		producer()
	})

	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertCalled,
		panicAssert:   assertNotCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnErrIs(errTestErrorInstance, func() {
			th.called()
		})

		producer()
	})
}

func Test_Handler_OnErrAs(t *testing.T) {
	oldTestErrorInstance := errTestErrorInstance
	defer func() { errTestErrorInstance = oldTestErrorInstance }()

	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertNotCalled,
		panicAssert:   assertNotCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnErrAs(func(err net.Error) {
			th.calledErr(err)
		})

		producer()
	})

	errTestErrorInstance = &net.OpError{Op: "op", Err: oldTestErrorInstance}

	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertCalledTestErr,
		panicAssert:   assertNotCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnErrAs(func(err net.Error) {
			th.calledErr(err)
		})

		producer()
	})

	errTestErrorInstance = errType{}

	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertCalledTestErr,
		panicAssert:   assertNotCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnErrAs(func(err errType) {
			th.calledErr(err)
		})

		producer()
	})

	errTestErrorInstance = &errType{}

	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertCalledTestErr,
		panicAssert:   assertNotCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnErrAs(func(err *errType) {
			th.calledErr(err)
		})

		producer()
	})

	errTestErrorInstance = &errPtrType{}

	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertCalledTestErr,
		panicAssert:   assertNotCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnErrAs(func(err *errPtrType) {
			th.calledErr(err)
		})

		producer()
	})
}

func Test_Handler_OnErrOrPanic(t *testing.T) {
	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertCalledTestErr,
		panicAssert:   assertCalledTestPanic,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnErrOrPanic(func(err error) {
			th.calledErr(err)
		})

		producer()
	})
}

func Test_Handler_OnPanic(t *testing.T) {
	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertNotCalled,
		panicAssert:   assertCalledTestPanic,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnPanic(func(panicObj interface{}) {
			th.calledPanic(panicObj)
		})

		producer()
	})
}

func Test_Handler_OnAnyPanic(t *testing.T) {
	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertNotCalled,
		panicAssert:   assertCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnAnyPanic(func() {
			th.called()
		})

		producer()
	})
}

func Test_Handler_OnAnyErr(t *testing.T) {
	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertCalled,
		panicAssert:   assertNotCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnAnyErr(func() {
			th.called()
		})

		producer()
	})
}

func Test_Handler_OnAnyErrOrPanic(t *testing.T) {
	doTest(t, testConfig{
		successAssert: assertNotCalled,
		errorAssert:   assertCalled,
		panicAssert:   assertCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnAnyErrOrPanic(func() {
			th.called()
		})

		producer()
	})
}

func Test_Handler_OnSuccess(t *testing.T) {
	doTest(t, testConfig{
		successAssert: assertCalled,
		errorAssert:   assertNotCalled,
		panicAssert:   assertNotCalled,
	}, func(th *testHelper, producer func()) {
		defer IfError().ThenIgnore()

		defer Handle().OnSuccess(func() {
			th.called()
		})

		producer()
	})
}

func Test_Handler_OnErr_Wrapper(t *testing.T) {
	var resultErr error
	fn := func() {
		defer IfError().ThenIgnore()

		defer Handle().OnErr(func(err error) {
			resultErr = err
		})

		With(WrapperFmtErrorw("wrapper")).CheckErr(fmt.Errorf("error"))
	}

	fn()
	assert.EqualError(t, resultErr, "wrapper: error")
}

func Test_IsErr(t *testing.T) {
	assert.True(t, IsErr(fmt.Errorf("error")))
	assert.False(t, IsErr(PanicErr{PanicObj: "panic"}))
	assert.False(t, IsErr(nil))
}

func Test_IsSuccess(t *testing.T) {
	assert.False(t, IsSuccess(fmt.Errorf("error")))
	assert.False(t, IsSuccess(PanicErr{PanicObj: "panic"}))
	assert.True(t, IsSuccess(nil))
}

func Test_IsPanic(t *testing.T) {
	assert.False(t, IsPanic(fmt.Errorf("error")))
	assert.True(t, IsPanic(PanicErr{PanicObj: "panic"}))
	assert.False(t, IsPanic(nil))
}

func Test_GetPanic(t *testing.T) {
	var panicObj interface{}

	assert.False(t, GetPanic(fmt.Errorf("error"), &panicObj))
	assert.Nil(t, panicObj)

	assert.True(t, GetPanic(PanicErr{PanicObj: "panic"}, &panicObj))
	assert.Equal(t, "panic", panicObj)
}

func Test_verifyErrFnType(t *testing.T) {
	assert.PanicsWithError(t, "arg should be a function", func() {
		verifyErrFnType("arg", "string")
	})

	assert.PanicsWithError(t, "arg should have exactly 1 input argument", func() {
		verifyErrFnType("arg", func(err1, err2 error) {})
	})

	assert.PanicsWithError(t, "arg should have exactly 1 input argument", func() {
		verifyErrFnType("arg", func(errs ...error) {})
	})

	assert.PanicsWithError(t, "arg should have exactly 1 input argument", func() {
		verifyErrFnType("arg", func(err error, errs ...error) {})
	})

	assert.PanicsWithError(t, "arg should have exactly no output arguments", func() {
		verifyErrFnType("arg", func(err error) error { return nil })
	})

	assert.PanicsWithError(t, "arg first argument should be assignable to error interface", func() {
		verifyErrFnType("arg", func(err string) {})
	})

	assert.PanicsWithError(t, "arg first argument should be assignable to error interface", func() {
		verifyErrFnType("arg", func(err *error) {})
	})

	assert.PanicsWithError(t, "arg first argument should be assignable to error interface", func() {
		verifyErrFnType("arg", func(err errPtrType) {})
	})

	assert.NotPanics(t, func() {
		verifyErrFnType("arg", func(err error) {})
	})

	assert.NotPanics(t, func() {
		verifyErrFnType("arg", func(err net.Error) {})
	})

	assert.NotPanics(t, func() {
		verifyErrFnType("arg", func(err *errPtrType) {})
	})

	assert.NotPanics(t, func() {
		verifyErrFnType("arg", func(err errType) {})
	})

	assert.NotPanics(t, func() {
		verifyErrFnType("arg", func(err *errType) {})
	})
}

func Test_Handler_NestedCheck_Success_Success(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)

		defer Handle().Always(func() {

		})

		return nil
	}

	assert.Nil(t, fn())
}

func Test_Handler_NestedCheck_Success_Error(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)

		defer Handle().Always(func() {
			CheckErr(fmt.Errorf("error 1"))
		})

		return nil
	}

	assert.EqualError(t, fn(), "error 1")
}

func Test_Handler_NestedCheck_Success_Panic(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)

		defer Handle().Always(func() {
			panic("panic 1")
		})

		return nil
	}

	assert.PanicsWithValue(t, "panic 1", func() {
		fn()
	})
}

func Test_Handler_NestedCheck_Error_Success(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)

		defer Handle().Always(func() {

		})

		return CheckErr(fmt.Errorf("error 1"))
	}

	assert.EqualError(t, fn(), "error 1")
}

func Test_Handler_NestedCheck_Error_Error(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)

		defer Handle().Always(func() {
			CheckErr(fmt.Errorf("error 2"))
		})

		return CheckErr(fmt.Errorf("error 1"))
	}

	assert.EqualError(t, fn(), "error 1 (also: error 2)")
}

func Test_Handler_NestedCheck_Error_Panic(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)

		defer Handle().Always(func() {
			panic("panic 1")
		})

		return CheckErr(fmt.Errorf("error 1"))
	}

	assert.PanicsWithValue(t, "panic 1", func() {
		fn()
	})
}

func Test_Handler_NestedCheck_Panic_Success(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)

		defer Handle().Always(func() {

		})

		panic("panic 1")
	}

	assert.PanicsWithValue(t, "panic 1", func() {
		fn()
	})
}

func Test_Handler_NestedCheck_Panic_Error(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)

		defer Handle().Always(func() {
			CheckErr(fmt.Errorf("error 1"))
		})

		panic("panic 1")
	}

	assert.PanicsWithValue(t, "panic 1", func() {
		fn()
	})
}

func Test_Handler_NestedCheck_Panic_Panic(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)

		defer Handle().Always(func() {
			panic("panic 2")
		})

		panic("panic 1")
	}

	assert.PanicsWithValue(t, "panic 2", func() {
		fn()
	})
}

func Test_Handler_DoubleNestedCheck(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ReturnWrapped().ThenAssignTo(&err)

		defer Handle().Always(func() {
			func() {
				CheckErr(fmt.Errorf("error 1"))
			}()
		})

		return nil
	}

	assert.PanicsWithError(t, "errflow incorrect call sequence", func() {
		fn()
	})
}
