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

func (ig *InterfaceGenerator) GenerateCode() (string, error) {
	if ig.Name == "" {
		return "", errmsg.InterfaceNameIsEmptyError()
	}

	stmt := fmt.Sprintf("type %s interface {\n", ig.Name)
	for _, sig := range ig.FuncSignatures {
		signatureStr, err := sig.GenerateCode()
		if err != nil {
			return "", err
		}
		stmt += fmt.Sprintf("\t%s\n", signatureStr)
	}
	stmt += "}"

	return stmt, nil
}
