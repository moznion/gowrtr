package generator

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// AnonymousFuncSignatureGenerator represents a code generator for signature of anonymous func.
type AnonymousFuncSignatureGenerator struct {
	FuncParameters []*FuncParameter
	ReturnTypes    []string
}

// NewAnonymousFuncSignatureGenerator returns a new `AnonymousFuncSignatureGenerator`.
func NewAnonymousFuncSignatureGenerator() *AnonymousFuncSignatureGenerator {
	return &AnonymousFuncSignatureGenerator{}
}

// AddFuncParameters adds parameters of function to `AnonymousFuncSignatureGenerator`.
// This method returns a *new* `AnonymousFuncSignatureGenerator`; it means this method acts as immutable.
func (f *AnonymousFuncSignatureGenerator) AddFuncParameters(funcParameters ...*FuncParameter) *AnonymousFuncSignatureGenerator {
	return &AnonymousFuncSignatureGenerator{
		FuncParameters: append(f.FuncParameters, funcParameters...),
		ReturnTypes:    f.ReturnTypes,
	}
}

// AddReturnTypes adds return types of the function to `AnonymousFuncSignatureGenerator`.
// This method returns a *new* `AnonymousFuncSignatureGenerator`; it means this method acts as immutable.
func (f *AnonymousFuncSignatureGenerator) AddReturnTypes(returnTypes ...string) *AnonymousFuncSignatureGenerator {
	return &AnonymousFuncSignatureGenerator{
		FuncParameters: f.FuncParameters,
		ReturnTypes:    append(f.ReturnTypes, returnTypes...),
	}
}

// Generate generates a signature of the anonymous func as golang code.
func (f *AnonymousFuncSignatureGenerator) Generate(indentLevel int) (string, error) {
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
