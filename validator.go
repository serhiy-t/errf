package errflow

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var globalErrflowValidator validator = &noopValidator{}

func init() {
	if strings.HasSuffix(os.Args[0], ".test") || strings.HasSuffix(os.Args[0], ".test.exe") {
		SetStackTraceValidator()
	}
}

type restoreValidatorCloser struct {
	oldValidator validator
}

func (rvc *restoreValidatorCloser) Close() error {
	globalErrflowValidator = rvc.oldValidator
	return nil
}

func setValidator(v validator) io.Closer {
	oldValidator := globalErrflowValidator
	globalErrflowValidator = v
	return &restoreValidatorCloser{
		oldValidator: oldValidator,
	}
}

// SetNoopValidator sets no-op validator for errflow.
//
// Validator is used to validate that the library is used correctly,
// meaning each used is limited to a single function.
//
// This is a default mode for production, which doesn't compromise performance,
// but library can be misused in this mode.
//
// It returns io.Closer instance,
// which can be used to restore previous validator,
// but also can be ignored, if not needed.
func SetNoopValidator() io.Closer {
	return setValidator(&noopValidator{})
}

// SetStackTraceValidator sets a stack-trace based validator for errflow.
//
// Validator is used to validate that the library is used correctly,
// meaning each used is limited to a single function.
//
// This is a default mode for tests, which works in most cases, but
// has performance penalty and might return false positives in some cases.
//
// It returns io.Closer instance,
// which can be used to restore previous validator,
// but also can be ignored, if not needed.
func SetStackTraceValidator() io.Closer {
	return setValidator(&stackTraceValidator{})
}

type validator interface {
	enter()
	leave()
	validate()
}

type noopValidator struct {
}

func (v *noopValidator) enter()    {}
func (v *noopValidator) leave()    {}
func (v *noopValidator) validate() {}

type stackTraceValidator struct {
}

func (v *stackTraceValidator) enter() {
	getGoroutineErrflowStack().push()
}

func (v *stackTraceValidator) leave() {
	getGoroutineErrflowStack().pop()
}

func (v *stackTraceValidator) validate() {
	getGoroutineErrflowStack().validate()
}

type errflowStack struct {
	stack []string
}

func (s *errflowStack) push() {
	s.stack = append(s.stack, getCurrentCallerFn())
}

func (s *errflowStack) pop() {
	s.validate()
	s.stack = s.stack[:len(s.stack)-1]
	cleanupGoroutineErrflowStack()
}

func (s *errflowStack) validate() {
	if len(s.stack) == 0 || s.stack[len(s.stack)-1] != getCurrentCallerFn() {
		panic(fmt.Errorf("errflow incorrect call sequence"))
	}
}

func getCurrentCallerFn() string {
	depth := 1
	for {
		c, _, _, _ := runtime.Caller(depth)
		fn := runtime.FuncForPC(c).Name()
		if strings.HasPrefix(fn, "runtime") || strings.HasPrefix(fn, "testing") {
			depth++
			continue
		}
		if strings.Contains(fn, "errflow.ImplementCheck") {
			depth += 2
			continue
		}
		if strings.Contains(fn, "errflow.") {
			depth++
			continue
		}
		return fn
	}
}

func (s *errflowStack) empty() bool {
	return len(s.stack) == 0
}

var goroutineErrflowStackMap = make(map[int]*errflowStack)

func getGoroutineErrflowStack() *errflowStack {
	goID := goid()
	_, ok := goroutineErrflowStackMap[goID]
	if !ok {
		goroutineErrflowStackMap[goID] = &errflowStack{}
	}
	return goroutineErrflowStackMap[goID]
}

func cleanupGoroutineErrflowStack() {
	goID := goid()
	errflowStack, ok := goroutineErrflowStackMap[goID]
	if ok {
		if errflowStack.empty() {
			delete(goroutineErrflowStackMap, goID)
		}
	}
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
