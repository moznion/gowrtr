package generator

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// FuncParameter represents a parameter of the func.
type FuncParameter struct {
	Name string
	Type string
}

// FuncSignature represents a code generator for the signature of the func.
type FuncSignature struct {
	FuncName       string
	FuncParameters []*FuncParameter
	ReturnTypes    []string
}

// NewFuncParameter returns a new `FuncSignature`.
func NewFuncParameter(name string, typ string) *FuncParameter {
	return &FuncParameter{
		Name: name,
		Type: typ,
	}
}

// NewFuncSignature returns a new `FuncSignature`.
func NewFuncSignature(funcName string) *FuncSignature {
	return &FuncSignature{
		FuncName: funcName,
	}
}

// AddFuncParameters adds parameters of the func to `FuncSignature`.
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) AddFuncParameters(funcParameters ...*FuncParameter) *FuncSignature {
	return &FuncSignature{
		FuncName:       f.FuncName,
		FuncParameters: append(f.FuncParameters, funcParameters...),
		ReturnTypes:    f.ReturnTypes,
	}
}

// AddReturnTypes adds return types of the func to `FuncSignature`.
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) AddReturnTypes(returnTypes ...string) *FuncSignature {
	return &FuncSignature{
		FuncName:       f.FuncName,
		FuncParameters: f.FuncParameters,
		ReturnTypes:    append(f.ReturnTypes, returnTypes...),
	}
}

// Generate generates a signature of the func as golang code.
func (f *FuncSignature) Generate(indentLevel int) (string, error) {
	if f.FuncName == "" {
		return "", errmsg.FuncNameIsEmptyError()
	}

	stmt := f.FuncName + "("

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
