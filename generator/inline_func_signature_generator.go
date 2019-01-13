package gowrtr

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

type InlineFuncSignatureGenerator struct {
	FuncParameters []*FuncParameter
	ReturnTypes    []string
}

func NewInlineFuncSignatureGenerator() *InlineFuncSignatureGenerator {
	return &InlineFuncSignatureGenerator{}
}

func (f *InlineFuncSignatureGenerator) AddFuncParameters(funcParameters ...*FuncParameter) *InlineFuncSignatureGenerator {
	return &InlineFuncSignatureGenerator{
		FuncParameters: append(f.FuncParameters, funcParameters...),
		ReturnTypes:    f.ReturnTypes,
	}
}

func (f *InlineFuncSignatureGenerator) AddReturnTypes(returnTypes ...string) *InlineFuncSignatureGenerator {
	return &InlineFuncSignatureGenerator{
		FuncParameters: f.FuncParameters,
		ReturnTypes:    append(f.ReturnTypes, returnTypes...),
	}
}

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
