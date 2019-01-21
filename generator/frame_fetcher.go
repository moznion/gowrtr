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

		return fmt.Sprintf("%s:%d", file, line)
	}

	return ""
}
