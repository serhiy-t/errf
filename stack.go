package errf

import (
	"runtime/debug"
	"strings"
)

type ParsedStackItem struct {
	Fn  string
	Src string
}

type ParsedStack struct {
	Goroutine string
	Items     []ParsedStackItem
}

func (ps ParsedStack) String() string {
	var lines []string
	lines = append(lines, ps.Goroutine)
	for _, item := range ps.Items {
		lines = append(lines, item.Fn)
		lines = append(lines, item.Src)
	}
	return strings.Join(lines, "\n")
}

func getStringErrorStackTraceFn() func() string {
	result := getErrorStackTrace()
	return func() string {
		return result.String()
	}
}

func getErrorStackTrace() ParsedStack {
	var result ParsedStack
	lines := strings.Split(strings.TrimSpace(string(debug.Stack())), "\n")
	result.Goroutine = lines[0]
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
		result.Items = append(result.Items, ParsedStackItem{
			Fn:  lines[idx],
			Src: lines[idx+1],
		})
	}
	return result
}
