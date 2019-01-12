package gowrtr

import "strings"

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
	return "(" + strings.Join(fig.Parameters, ", ") + ")", nil
}
