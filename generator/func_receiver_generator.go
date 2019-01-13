package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// FuncReceiverGenerator represents a code generator for the receiver of the func.
type FuncReceiverGenerator struct {
	Name string
	Type string
}

// NewFuncReceiverGenerator returns a new `FuncReceiverGenerator`.
func NewFuncReceiverGenerator(name string, typ string) *FuncReceiverGenerator {
	return &FuncReceiverGenerator{
		Name: name,
		Type: typ,
	}
}

// Generate generates a receiver of the func as golang code.
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
