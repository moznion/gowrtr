package generator

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// InlineFuncSignatureGenerator represents a code generator for signature of inline func.
type InlineFuncSignatureGenerator struct {
	FuncParameters []*FuncParameter
	ReturnTypes    []string
}

// NewInlineFuncSignatureGenerator returns a new `InlineFuncSignatureGenerator`.
func NewInlineFuncSignatureGenerator() *InlineFuncSignatureGenerator {
	return &InlineFuncSignatureGenerator{}
}

// AddFuncParameters adds parameters of function to `InlineFuncSignatureGenerator`.
// This method returns a *new* `InlineFuncSignatureGenerator`; it means this method acts as immutable.
func (f *InlineFuncSignatureGenerator) AddFuncParameters(funcParameters ...*FuncParameter) *InlineFuncSignatureGenerator {
	return &InlineFuncSignatureGenerator{
		FuncParameters: append(f.FuncParameters, funcParameters...),
		ReturnTypes:    f.ReturnTypes,
	}
}

// AddReturnTypes adds return types of the function to `InlineFuncSignatureGenerator`.
// This method returns a *new* `InlineFuncSignatureGenerator`; it means this method acts as immutable.
func (f *InlineFuncSignatureGenerator) AddReturnTypes(returnTypes ...string) *InlineFuncSignatureGenerator {
	return &InlineFuncSignatureGenerator{
		FuncParameters: f.FuncParameters,
		ReturnTypes:    append(f.ReturnTypes, returnTypes...),
	}
}

// Generate generates a signature of the inline func as golang code.
func (f *InlineFuncSignatureGenerator) Generate(indentLevel int) (string, error) {
	stmt := "("

	typeExisted := true
	params := make([]string, len(f.FuncParameters))
	for i, param := range f.FuncParameters {
		if param.Name == "" {
			return "", errmsg.FuncParameterNameIsEmptyErr()
		}

		paramSet := param.Name
		typeExisted = param.Type != ""
		if typeExisted {
			paramSet += " " + param.Type
		}
		params[i] = paramSet
	}

	if !typeExisted {
		return "", errmsg.LastFuncParameterTypeIsEmptyErr()
	}

	stmt += strings.Join(params, ", ") + ")"

	returnTypes := f.ReturnTypes
	switch len(returnTypes) {
	case 0:
		// NOP
	case 1:
		stmt += " " + returnTypes[0]
	default:
		stmt += " (" + strings.Join(returnTypes, ", ") + ")"
	}
	return stmt, nil
}
