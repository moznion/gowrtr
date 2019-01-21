package generator

import (
	"fmt"
	"runtime"
	"strings"
)

func fetchClientCallerLine(skip ...int) string {
	s := 2
	if len(skip) > 0 {
		s = skip[0]
	}

	caller := ""
	for {
		pc, file, line, ok := runtime.Caller(s)
		f := runtime.FuncForPC(pc)
		if strings.Contains(f.Name(), "github.com/moznion/gowrtr") {
			s++
			continue
		}

		if !ok {
			break
		}

		caller = fmt.Sprintf("%s:%d", file, line)
		break
	}

	return caller
}

func fetchClientCallerLineAsSlice(size int, skip ...int) []string {
	s := 3
	if len(skip) > 0 {
		s = skip[0]
	}

	caller := fetchClientCallerLine(s)
	callers := make([]string, size)
	for i := 0; i < size; i++ {
		callers[i] = caller
	}
	return callers
}
