package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// FuncReceiver represents a code generator for the receiver of the func.
type FuncReceiver struct {
	name                       string
	typ                        string
	caller                     string
	genericsTypeParameterNames TypeParameterNames
}

// NewFuncReceiver returns a new `FuncReceiver`.
func NewFuncReceiver(name string, typ string) *FuncReceiver {
	return NewFuncReceiverWithGenerics(name, typ, TypeParameterNames{})
}

func NewFuncReceiverWithGenerics(name string, typ string, genericsTypeParameterNames TypeParameterNames) *FuncReceiver {
	return &FuncReceiver{
		name:                       name,
		typ:                        typ,
		caller:                     fetchClientCallerLine(),
		genericsTypeParameterNames: genericsTypeParameterNames,
	}
}

// Generate generates a receiver of the func as golang code.
func (f *FuncReceiver) Generate(indentLevel int) (string, error) {
	name := f.name
	typ := f.typ

	if typ == "" && name == "" {
		return "", nil
	}

	if name == "" {
		return "", errmsg.FuncReceiverNameIsEmptyError(f.caller)
	}

	if typ == "" {
		return "", errmsg.FuncReceiverTypeIsEmptyError(f.caller)
	}

	genericsTypeParam := ""
	if len(f.genericsTypeParameterNames) > 0 {
		genericsTypeParam, _ = f.genericsTypeParameterNames.Generate(0)
	}

	return fmt.Sprintf("(%s %s%s)", name, typ, genericsTypeParam), nil
}
