package errf

import (
	"runtime/debug"
	"strings"
)

type parsedStackItem struct {
	fn  string
	src string
}

type parsedStack struct {
	goroutine string
	items     []parsedStackItem
}

func (ps parsedStack) String() string {
	var lines []string
	lines = append(lines, ps.goroutine)
	for _, item := range ps.items {
		lines = append(lines, item.fn)
		lines = append(lines, item.src)
	}
	return strings.Join(lines, "\n")
}

func getStringErrorStackTraceFn() func() string {
	debugStack := debug.Stack()
	return func() string {
		return parseErrorStackTrace(string(debugStack)).String()
	}
}

func parseErrorStackTrace(debugStack string) parsedStack {
	var result parsedStack
	lines := strings.Split(strings.TrimSpace(debugStack), "\n")
	result.goroutine = lines[0]
	lines = lines[1:]
	idx := 0
	for ; idx < len(lines)-1; idx += 2 {
		if strings.HasPrefix(lines[idx], "runtime") ||
			strings.HasPrefix(lines[idx], "testing") ||
			strings.HasPrefix(lines[idx], "panic(") {
			continue
		}
		if strings.Contains(lines[idx], "/errf.(*Errflow).ImplementCheck(") {
			idx += 2
			continue
		}
		if strings.Contains(lines[idx], "/errf.") && !strings.Contains(lines[idx+1], "_test.go") {
			continue
		}
		break
	}
	if idx >= len(lines)-1 {
		idx = 0
	}
	for ; idx < len(lines)-1; idx += 2 {
		result.items = append(result.items, parsedStackItem{
			fn:  lines[idx],
			src: lines[idx+1],
		})
	}
	return result
}

func getErrorStackTrace() parsedStack {
	return parseErrorStackTrace(string(debug.Stack()))
}
