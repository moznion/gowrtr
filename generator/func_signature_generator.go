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

// FuncSignatureGenerator represents a code generator for the signature of the func.
type FuncSignatureGenerator struct {
	FuncName       string
	FuncParameters []*FuncParameter
	ReturnTypes    []string
}

// NewFuncParameter returns a new `FuncSignatureGenerator`.
func NewFuncParameter(name string, typ string) *FuncParameter {
	return &FuncParameter{
		Name: name,
		Type: typ,
	}
}

// NewFuncSignatureGenerator returns a new `FuncSignatureGenerator`.
func NewFuncSignatureGenerator(funcName string) *FuncSignatureGenerator {
	return &FuncSignatureGenerator{
		FuncName: funcName,
	}
}

// AddFuncParameters adds parameters of the func to `FuncSignatureGenerator`.
// This method returns a *new* `FuncSignatureGenerator`; it means this method acts as immutable.
func (f *FuncSignatureGenerator) AddFuncParameters(funcParameters ...*FuncParameter) *FuncSignatureGenerator {
	return &FuncSignatureGenerator{
		FuncName:       f.FuncName,
		FuncParameters: append(f.FuncParameters, funcParameters...),
		ReturnTypes:    f.ReturnTypes,
	}
}

// AddReturnTypes adds return types of the func to `FuncSignatureGenerator`.
// This method returns a *new* `FuncSignatureGenerator`; it means this method acts as immutable.
func (f *FuncSignatureGenerator) AddReturnTypes(returnTypes ...string) *FuncSignatureGenerator {
	return &FuncSignatureGenerator{
		FuncName:       f.FuncName,
		FuncParameters: f.FuncParameters,
		ReturnTypes:    append(f.ReturnTypes, returnTypes...),
	}
}

// Generate generates a signature of the func as golang's code.
func (f *FuncSignatureGenerator) Generate(indentLevel int) (string, error) {
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
