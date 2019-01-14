package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// Interface represents a code generator for `interface` block.
type Interface struct {
	Name           string
	FuncSignatures []*FuncSignature
}

// NewInterface returns a new `Interface`.
func NewInterface(name string, funcSignatures ...*FuncSignature) *Interface {
	return &Interface{
		Name:           name,
		FuncSignatures: funcSignatures,
	}
}

// AddFuncSignatures adds signatures of the func to `Interface`.
// This method returns a *new* `Interface`; it means this method acts as immutable.
func (ig *Interface) AddFuncSignatures(sig ...*FuncSignature) *Interface {
	return &Interface{
		Name:           ig.Name,
		FuncSignatures: append(ig.FuncSignatures, sig...),
	}
}

// Generate generates `interface` block as golang code.
func (ig *Interface) Generate(indentLevel int) (string, error) {
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
