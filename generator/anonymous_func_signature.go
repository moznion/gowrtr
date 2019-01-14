package generator

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// AnonymousFuncSignature represents a code generator for signature of anonymous func.
type AnonymousFuncSignature struct {
	FuncParameters []*FuncParameter
	ReturnTypes    []string
}

// NewAnonymousFuncSignature returns a new `AnonymousFuncSignature`.
func NewAnonymousFuncSignature() *AnonymousFuncSignature {
	return &AnonymousFuncSignature{}
}

// AddFuncParameters adds parameters of function to `AnonymousFuncSignature`.
// This method returns a *new* `AnonymousFuncSignature`; it means this method acts as immutable.
func (f *AnonymousFuncSignature) AddFuncParameters(funcParameters ...*FuncParameter) *AnonymousFuncSignature {
	return &AnonymousFuncSignature{
		FuncParameters: append(f.FuncParameters, funcParameters...),
		ReturnTypes:    f.ReturnTypes,
	}
}

// AddReturnTypes adds return types of the function to `AnonymousFuncSignature`.
// This method returns a *new* `AnonymousFuncSignature`; it means this method acts as immutable.
func (f *AnonymousFuncSignature) AddReturnTypes(returnTypes ...string) *AnonymousFuncSignature {
	return &AnonymousFuncSignature{
		FuncParameters: f.FuncParameters,
		ReturnTypes:    append(f.ReturnTypes, returnTypes...),
	}
}

// Generate generates a signature of the anonymous func as golang code.
func (f *AnonymousFuncSignature) Generate(indentLevel int) (string, error) {
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
