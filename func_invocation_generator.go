package gowrtr

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

type FuncInvocationGenerator struct {
	Parameters []string
}

func NewFuncInvocationGenerator(parameters ...string) *FuncInvocationGenerator {
	return &FuncInvocationGenerator{
		Parameters: parameters,
	}
}

func (fig *FuncInvocationGenerator) AddParameters(parameters ...string) *FuncInvocationGenerator {
	return NewFuncInvocationGenerator(append(fig.Parameters, parameters...)...)
}

func (fig *FuncInvocationGenerator) Generate(indentLevel int) (string, error) {
	for _, param := range fig.Parameters {
		if param == "" {
			return "", errmsg.FuncInvocationParameterIsEmptyError()
		}
	}

	return "(" + strings.Join(fig.Parameters, ", ") + ")", nil
}
