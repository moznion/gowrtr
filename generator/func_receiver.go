package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// FuncReceiver represents a code generator for the receiver of the func.
type FuncReceiver struct {
	Name string
	Type string
}

// NewFuncReceiver returns a new `FuncReceiver`.
func NewFuncReceiver(name string, typ string) *FuncReceiver {
	return &FuncReceiver{
		Name: name,
		Type: typ,
	}
}

// Generate generates a receiver of the func as golang code.
func (f *FuncReceiver) Generate(indentLevel int) (string, error) {
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
