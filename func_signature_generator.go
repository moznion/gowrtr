package gowrtr

import (
	"strings"

	"github.com/moznion/gowrtr/errmsg"
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

func NewFuncSignatureGenerator(funcName string, funcParameters []*FuncParameter, returnTypes []string) *FuncSignatureGenerator {
	return &FuncSignatureGenerator{
		FuncName:       funcName,
		FuncParameters: funcParameters,
		ReturnTypes:    returnTypes,
	}
}

func (f *FuncSignatureGenerator) AddFuncParameters(funcParameters ...*FuncParameter) *FuncSignatureGenerator {
	f.FuncParameters = append(f.FuncParameters, funcParameters...)
	return f
}

func (f *FuncSignatureGenerator) AddReturnTypes(returnTypes ...string) *FuncSignatureGenerator {
	f.ReturnTypes = append(f.ReturnTypes, returnTypes...)
	return f
}

func (f *FuncSignatureGenerator) GenerateCode(indentLevel int) (string, error) {
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
