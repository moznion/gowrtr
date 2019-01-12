package gowrtr

import (
	"fmt"

	"github.com/moznion/gowrtr/errmsg"
)

type InterfaceGenerator struct {
	Name           string
	FuncSignatures []*FuncSignatureGenerator
}

func NewInterfaceGenerator(name string, funcSignatures ...*FuncSignatureGenerator) *InterfaceGenerator {
	return &InterfaceGenerator{
		Name:           name,
		FuncSignatures: funcSignatures,
	}
}

func (ig *InterfaceGenerator) AddFuncSignature(sig *FuncSignatureGenerator) *InterfaceGenerator {
	return &InterfaceGenerator{
		Name:           ig.Name,
		FuncSignatures: append(ig.FuncSignatures, sig),
	}
}

func (ig *InterfaceGenerator) Generate(indentLevel int) (string, error) {
	if ig.Name == "" {
		return "", errmsg.InterfaceNameIsEmptyError()
	}

	indent := buildIndent(indentLevel)

	nextIndentLevel := indentLevel + 1
	stmt := fmt.Sprintf("%stype %s interface {\n", indent, ig.Name)
	for _, sig := range ig.FuncSignatures {
		signatureStr, err := sig.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += fmt.Sprintf("%s\t%s\n", indent, signatureStr)
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
