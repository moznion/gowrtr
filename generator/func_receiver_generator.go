package gowrtr

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

type FuncReceiverGenerator struct {
	Name string
	Type string
}

func NewFuncReceiverGenerator(name string, typ string) *FuncReceiverGenerator {
	return &FuncReceiverGenerator{
		Name: name,
		Type: typ,
	}
}

func (f *FuncReceiverGenerator) Generate(indentLevel int) (string, error) {
	name := f.Name
	typ := f.Type

	if typ == "" && name == "" {
		return "", nil
	}

	if name == "" {
		return "", errmsg.FuncReceiverNameIsEmptyError()
	}

	if typ == "" {
		return "", errmsg.FuncReceiverTypeIsEmptyError()
	}

	return fmt.Sprintf("(%s %s)", name, typ), nil
}
