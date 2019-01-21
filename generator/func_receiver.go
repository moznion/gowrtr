package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// FuncReceiver represents a code generator for the receiver of the func.
type FuncReceiver struct {
	name   string
	typ    string
	caller string
}

// NewFuncReceiver returns a new `FuncReceiver`.
func NewFuncReceiver(name string, typ string) *FuncReceiver {
	return &FuncReceiver{
		name:   name,
		typ:    typ,
		caller: fetchClientCallerLine(),
	}
}

// Generate generates a receiver of the func as golang code.
func (f *FuncReceiver) Generate(indentLevel int) (string, error) {
	name := f.name
	typ := f.typ

	if typ == "" && name == "" {
		return "", nil
	}

	if name == "" {
		return "", errmsg.FuncReceiverNameIsEmptyError(f.caller)
	}

	if typ == "" {
		return "", errmsg.FuncReceiverTypeIsEmptyError(f.caller)
	}

	return fmt.Sprintf("(%s %s)", name, typ), nil
}
