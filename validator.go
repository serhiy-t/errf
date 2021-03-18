package errf

import (
	"fmt"
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

type restoreValidatorRestorer struct {
	oldValidator validator
}

func (rvc *restoreValidatorRestorer) ThenRestore() {
	globalErrflowValidator = rvc.oldValidator
}

func setValidator(v validator) DeferRestorer {
	oldValidator := globalErrflowValidator
	globalErrflowValidator = v
	return &restoreValidatorRestorer{
		oldValidator: oldValidator,
	}
}

// SetNoopValidator sets no-op validator for errflow.
//
// Validator is used to validate that the library is used correctly,
// meaning each usage is limited to a single function.
//
// This is a default mode for production, which doesn't compromise performance,
// but library can be misused in this mode.
//
// It returns errf.DeferRestorer instance,
// which can be used to restore previous validator, if needed.
func SetNoopValidator() DeferRestorer {
	return setValidator(&noopValidator{})
}

// SetStackTraceValidator sets a stack-trace based validator for errflow.
//
// Validator is used to validate that the library is used correctly,
// meaning each usage is limited to a single function.
//
// This is a default mode for tests, which works in most cases, but
// has performance penalty and might return false positives in some cases.
//
// It returns errf.DeferRestorer instance,
// which can be used to restore previous validator, if needed.
func SetStackTraceValidator() DeferRestorer {
	return setValidator(&stackTraceValidator{})
}

type validator interface {
	enter()
	leave()
	markPanic()
	validate()
	custom(func())
}

type noopValidator struct {
}

func (v *noopValidator) enter()        {}
func (v *noopValidator) leave()        {}
func (v *noopValidator) markPanic()    {}
func (v *noopValidator) validate()     {}
func (v *noopValidator) custom(func()) {}

type stackTraceValidator struct {
}

func (v *stackTraceValidator) enter() {
	getGoroutineErrflowStack().push()
}

func (v *stackTraceValidator) leave() {
	getGoroutineErrflowStack().pop()
}

func (v *stackTraceValidator) markPanic() {
	getGoroutineErrflowStack().markPanic = true
}

func (v *stackTraceValidator) validate() {
	getGoroutineErrflowStack().validate()
}

func (v *stackTraceValidator) custom(fn func()) {
	fn()
}

type errflowStack struct {
	stack     []string
	markPanic bool
}

func (s *errflowStack) push() {
	s.stack = append(s.stack, getCurrentCallerFn())
}

func (s *errflowStack) pop() {
	s.validate()
	s.stack = s.stack[:len(s.stack)-1]
	s.markPanic = false
	cleanupGoroutineErrflowStack()
}

func (s *errflowStack) validate() {
	if s.markPanic {
		return
	}
	currentCallerFn := getCurrentCallerFn()
	if len(s.stack) == 0 || s.stack[len(s.stack)-1] != currentCallerFn {
		panic(fmt.Errorf("errflow incorrect call sequence"))
	}
}

func getCurrentCallerFn() string {
	parsedStack := getErrorStackTrace()
	if len(parsedStack.items) == 0 {
		return "<unknown>"
	}
	fn := parsedStack.items[0].fn
	pIdx := strings.Index(fn, "(")
	if pIdx != -1 {
		fn = fn[:pIdx+1]
	}
	return fn
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
