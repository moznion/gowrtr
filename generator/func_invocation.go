package generator

import (
	"fmt"
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// FuncInvocation represents a code generator for func invocation.
type FuncInvocation struct {
	parameters    []string
	callers       []string
	genericsTypes []string
}

// NewFuncInvocation returns a new `FuncInvocation`.
func NewFuncInvocation(parameters ...string) *FuncInvocation {
	return &FuncInvocation{
		parameters:    parameters,
		callers:       fetchClientCallerLineAsSlice(len(parameters)),
		genericsTypes: []string{},
	}
}

// AddParameters adds parameters of func invocation to `FuncInvocation`. This does *not* set, just add.
// This method returns a *new* `FuncInvocation`; it means this method acts as immutable.
func (fig *FuncInvocation) AddParameters(parameters ...string) *FuncInvocation {
	return &FuncInvocation{
		parameters:    append(fig.parameters, parameters...),
		callers:       append(fig.callers, fetchClientCallerLineAsSlice(len(parameters))...),
		genericsTypes: fig.genericsTypes,
	}
}

// Parameters sets parameters of func invocation to `FuncInvocation`. This does *not* add, just set.
// This method returns a *new* `FuncInvocation`; it means this method acts as immutable.
func (fig *FuncInvocation) Parameters(parameters ...string) *FuncInvocation {
	return &FuncInvocation{
		parameters:    parameters,
		callers:       fetchClientCallerLineAsSlice(len(parameters)),
		genericsTypes: fig.genericsTypes,
	}
}

// GenericsTypes makes a new FuncInvocation value that is based on the receiver value with the given generics type names.
func (fig *FuncInvocation) GenericsTypes(typeNames ...string) *FuncInvocation {
	return &FuncInvocation{
		parameters:    fig.parameters,
		callers:       fig.callers,
		genericsTypes: typeNames,
	}
}

// Generate generates the func invocation as golang code.
func (fig *FuncInvocation) Generate(indentLevel int) (string, error) {
	for i, param := range fig.parameters {
		if param == "" {
			return "", errmsg.FuncInvocationParameterIsEmptyError(fig.callers[i])
		}
	}

	generics := ""
	if len(fig.genericsTypes) > 0 {
		generics = fmt.Sprintf("[%s]", strings.Join(fig.genericsTypes, ", "))
	}

	return generics + "(" + strings.Join(fig.parameters, ", ") + ")", nil
}
