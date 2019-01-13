package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// InterfaceGenerator represents a code generator for `interface` block.
type InterfaceGenerator struct {
	Name           string
	FuncSignatures []*FuncSignatureGenerator
}

// NewInterfaceGenerator returns a new `InterfaceGenerator`.
func NewInterfaceGenerator(name string, funcSignatures ...*FuncSignatureGenerator) *InterfaceGenerator {
	return &InterfaceGenerator{
		Name:           name,
		FuncSignatures: funcSignatures,
	}
}

// AddFuncSignatures adds signatures of the func to `InterfaceGenerator`.
// This method returns a *new* `InterfaceGenerator`; it means this method acts as immutable.
func (ig *InterfaceGenerator) AddFuncSignatures(sig ...*FuncSignatureGenerator) *InterfaceGenerator {
	return &InterfaceGenerator{
		Name:           ig.Name,
		FuncSignatures: append(ig.FuncSignatures, sig...),
	}
}

// Generate generates `interface` block as golang code.
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
