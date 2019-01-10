package gowrtr

import (
	"fmt"

	"github.com/moznion/gowrtr/errmsg"
)

type FuncReceiver struct {
	Name string
	Type string
}

func NewFuncReceiver(name string, typ string) *FuncReceiver {
	return &FuncReceiver{
		Name: name,
		Type: typ,
	}
}

func (f *FuncReceiver) GenerateCode() (string, error) {
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
