package gowrtr

import (
	"fmt"

	"github.com/moznion/gowrtr/errmsg"
)

type Interface struct {
	Name           string
	FuncSignatures []*FuncSignature
}

func NewInterface(name string, funcSignatures []*FuncSignature) *Interface {
	return &Interface{
		Name:           name,
		FuncSignatures: funcSignatures,
	}
}

func (in *Interface) GenerateCode() (string, error) {
	if in.Name == "" {
		return "", errmsg.InterfaceNameIsEmptyError()
	}

	stmt := fmt.Sprintf("type %s interface {\n", in.Name)
	for _, sig := range in.FuncSignatures {
		signatureStr, err := sig.GenerateCode()
		if err != nil {
			return "", err
		}
		stmt += fmt.Sprintf("\t%s\n", signatureStr)
	}
	stmt += "}"

	return stmt, nil
}
