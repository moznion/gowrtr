package generator

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// FuncInvocation represents a code generator for func invocation.
type FuncInvocation struct {
	Parameters []string
}

// NewFuncInvocation returns a new `FuncInvocation`.
func NewFuncInvocation(parameters ...string) *FuncInvocation {
	return &FuncInvocation{
		Parameters: parameters,
	}
}

// AddParameters adds parameters of func invocation to `FuncInvocation`.
// This method returns a *new* `FuncInvocation`; it means this method acts as immutable.
func (fig *FuncInvocation) AddParameters(parameters ...string) *FuncInvocation {
	return NewFuncInvocation(append(fig.Parameters, parameters...)...)
}

// Generate generates the func invocation as golang code.
func (fig *FuncInvocation) Generate(indentLevel int) (string, error) {
	for _, param := range fig.Parameters {
		if param == "" {
			return "", errmsg.FuncInvocationParameterIsEmptyError()
		}
	}

	return "(" + strings.Join(fig.Parameters, ", ") + ")", nil
}
