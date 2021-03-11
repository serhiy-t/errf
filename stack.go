package errflow

import (
	"runtime/debug"
	"strings"
)

func getErrorStackTrace() string {
	lines := strings.Split(string(debug.Stack()), "\n")
	firstLine := lines[0]
	lines = lines[1:]
	idx := 0
	for ; idx < len(lines)-1; idx += 2 {
		if strings.HasPrefix(lines[idx], "runtime") || strings.HasPrefix(lines[idx], "testing") {
			continue
		}
		if strings.Contains(lines[idx], "/errflow.ImplementTry") {
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
	return strings.Join(append([]string{firstLine}, lines[idx:]...), "\n")
}
