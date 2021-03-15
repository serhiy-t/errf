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
	result := getErrorStackTrace()
	return func() string {
		return result.String()
	}
}

func getErrorStackTrace() parsedStack {
	var result parsedStack
	lines := strings.Split(strings.TrimSpace(string(debug.Stack())), "\n")
	result.goroutine = lines[0]
	lines = lines[1:]
	idx := 0
	for ; idx < len(lines)-1; idx += 2 {
		if strings.HasPrefix(lines[idx], "runtime") ||
			strings.HasPrefix(lines[idx], "testing") ||
			strings.HasPrefix(lines[idx], "panic(") {
			continue
		}
		if strings.Contains(lines[idx], "/errflow.(*Errflow).ImplementTry(") {
			idx += 2
			continue
		}
		if strings.Contains(lines[idx], "/errflow.") && !strings.Contains(lines[idx+1], "_test.go") {
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
