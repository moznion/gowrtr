package gowrtr

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

type FuncParameter struct {
	Name string
	Type string
}

type FuncSignatureGenerator struct {
	FuncName       string
	FuncParameters []*FuncParameter
	ReturnTypes    []string
}

func NewFuncParameter(name string, typ string) *FuncParameter {
	return &FuncParameter{
		Name: name,
		Type: typ,
	}
}

func NewFuncSignatureGenerator(funcName string) *FuncSignatureGenerator {
	return &FuncSignatureGenerator{
		FuncName: funcName,
	}
}

func (f *FuncSignatureGenerator) AddFuncParameters(funcParameters ...*FuncParameter) *FuncSignatureGenerator {
	return &FuncSignatureGenerator{
		FuncName:       f.FuncName,
		FuncParameters: append(f.FuncParameters, funcParameters...),
		ReturnTypes:    f.ReturnTypes,
	}
}

func (f *FuncSignatureGenerator) AddReturnTypes(returnTypes ...string) *FuncSignatureGenerator {
	return &FuncSignatureGenerator{
		FuncName:       f.FuncName,
		FuncParameters: f.FuncParameters,
		ReturnTypes:    append(f.ReturnTypes, returnTypes...),
	}
}

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
