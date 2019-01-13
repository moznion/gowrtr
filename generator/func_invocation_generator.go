package generator

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// FuncInvocationGenerator represents a code generator for func invocation.
type FuncInvocationGenerator struct {
	Parameters []string
}

// NewFuncInvocationGenerator returns a new `FuncInvocationGenerator`.
func NewFuncInvocationGenerator(parameters ...string) *FuncInvocationGenerator {
	return &FuncInvocationGenerator{
		Parameters: parameters,
	}
}

// AddParameters adds parameters of func invocation.
func (fig *FuncInvocationGenerator) AddParameters(parameters ...string) *FuncInvocationGenerator {
	return NewFuncInvocationGenerator(append(fig.Parameters, parameters...)...)
}

// Generate generates the func invocation as golang's code.
func (fig *FuncInvocationGenerator) Generate(indentLevel int) (string, error) {
	for _, param := range fig.Parameters {
		if param == "" {
			return "", errmsg.FuncInvocationParameterIsEmptyError()
		}
	}

	return "(" + strings.Join(fig.Parameters, ", ") + ")", nil
}
