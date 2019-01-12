package gowrtr

import (
	"fmt"

	"github.com/moznion/gowrtr/errmsg"
)

type InterfaceGenerator struct {
	Name           string
	FuncSignatures []*FuncSignature
}

func NewInterfaceGenerator(name string, funcSignatures ...*FuncSignature) *InterfaceGenerator {
	return &InterfaceGenerator{
		Name:           name,
		FuncSignatures: funcSignatures,
	}
}

func (ig *InterfaceGenerator) AddFuncSignature(sig *FuncSignature) *InterfaceGenerator {
	ig.FuncSignatures = append(ig.FuncSignatures, sig)
	return ig
}

func (ig *InterfaceGenerator) GenerateCode(indentLevel int) (string, error) {
	if ig.Name == "" {
		return "", errmsg.InterfaceNameIsEmptyError()
	}

	indent := buildIndent(indentLevel)

	stmt := fmt.Sprintf("%stype %s interface {\n", indent, ig.Name)
	for _, sig := range ig.FuncSignatures {
		signatureStr, err := sig.GenerateCode()
		if err != nil {
			return "", err
		}
		stmt += fmt.Sprintf("%s\t%s\n", indent, signatureStr)
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
