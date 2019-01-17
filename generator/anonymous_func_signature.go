package generator

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// AnonymousFuncSignature represents a code generator for signature of anonymous func.
type AnonymousFuncSignature struct {
	funcParameters []*FuncParameter
	returnTypes    []string
}

// NewAnonymousFuncSignature returns a new `AnonymousFuncSignature`.
func NewAnonymousFuncSignature() *AnonymousFuncSignature {
	return &AnonymousFuncSignature{}
}

// AddParameters adds parameters of function to `AnonymousFuncSignature`. This does "not" set, just add.
// This method returns a *new* `AnonymousFuncSignature`; it means this method acts as immutable.
func (f *AnonymousFuncSignature) AddParameters(funcParameters ...*FuncParameter) *AnonymousFuncSignature {
	return &AnonymousFuncSignature{
		funcParameters: append(f.funcParameters, funcParameters...),
		returnTypes:    f.returnTypes,
	}
}

// Parameters sets parameters of function to `AnonymousFuncSignature`. This does "not" add, just set.
// This method returns a *new* `AnonymousFuncSignature`; it means this method acts as immutable.
func (f *AnonymousFuncSignature) Parameters(funcParameters ...*FuncParameter) *AnonymousFuncSignature {
	return &AnonymousFuncSignature{
		funcParameters: funcParameters,
		returnTypes:    f.returnTypes,
	}
}

// AddReturnTypes adds return types of the function to `AnonymousFuncSignature`. This does "not" set, just add.
// This method returns a *new* `AnonymousFuncSignature`; it means this method acts as immutable.
func (f *AnonymousFuncSignature) AddReturnTypes(returnTypes ...string) *AnonymousFuncSignature {
	return &AnonymousFuncSignature{
		funcParameters: f.funcParameters,
		returnTypes:    append(f.returnTypes, returnTypes...),
	}
}

// ReturnTypes sets return types of the function to `AnonymousFuncSignature`. This does "not" add, just set.
// This method returns a *new* `AnonymousFuncSignature`; it means this method acts as immutable.
func (f *AnonymousFuncSignature) ReturnTypes(returnTypes ...string) *AnonymousFuncSignature {
	return &AnonymousFuncSignature{
		funcParameters: f.funcParameters,
		returnTypes:    returnTypes,
	}
}

// Generate generates a signature of the anonymous func as golang code.
func (f *AnonymousFuncSignature) Generate(indentLevel int) (string, error) {
	stmt := "("

	typeExisted := true
	params := make([]string, len(f.funcParameters))
	for i, param := range f.funcParameters {
		if param.name == "" {
			return "", errmsg.FuncParameterNameIsEmptyErr()
		}

		paramSet := param.name
		typeExisted = param.typ != ""
		if typeExisted {
			paramSet += " " + param.typ
		}
		params[i] = paramSet
	}

	if !typeExisted {
		return "", errmsg.LastFuncParameterTypeIsEmptyErr()
	}

	stmt += strings.Join(params, ", ") + ")"

	returnTypes := f.returnTypes
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
