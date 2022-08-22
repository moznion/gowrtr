package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// Interface represents a code generator for `interface` block.
type Interface struct {
	name           string
	funcSignatures []*FuncSignature
	caller         string
	typeParameters TypeParameters
}

// NewInterface returns a new `Interface`.
func NewInterface(name string, funcSignatures ...*FuncSignature) *Interface {
	return &Interface{
		name:           name,
		funcSignatures: funcSignatures,
		caller:         fetchClientCallerLine(),
	}
}

// AddSignatures adds signatures of the func to `Interface`. This does *not* set, just add.
// This method returns a *new* `Interface`; it means this method acts as immutable.
func (ig *Interface) AddSignatures(sig ...*FuncSignature) *Interface {
	return &Interface{
		name:           ig.name,
		funcSignatures: append(ig.funcSignatures, sig...),
		caller:         ig.caller,
	}
}

// Signatures sets signatures of the func to `Interface`. This does *not* add, just set.
// This method returns a *new* `Interface`; it means this method acts as immutable.
func (ig *Interface) Signatures(sig ...*FuncSignature) *Interface {
	return &Interface{
		name:           ig.name,
		funcSignatures: sig,
		caller:         ig.caller,
	}
}

// TypeParameters makes a new Interface value based on the receiver value with the given generics TypeParameters.
func (ig *Interface) TypeParameters(typeParameters TypeParameters) *Interface {
	return &Interface{
		name:           ig.name,
		funcSignatures: ig.funcSignatures,
		caller:         ig.caller,
		typeParameters: typeParameters,
	}
}

// Generate generates `interface` block as golang code.
func (ig *Interface) Generate(indentLevel int) (string, error) {
	if ig.name == "" {
		return "", errmsg.InterfaceNameIsEmptyError(ig.caller)
	}

	indent := BuildIndent(indentLevel)

	typeParamStmt := ""
	if len(ig.typeParameters) > 0 {
		var err error
		typeParamStmt, err = ig.typeParameters.Generate(0)
		if err != nil {
			return "", err
		}
	}

	nextIndentLevel := indentLevel + 1
	stmt := fmt.Sprintf("%stype %s%s interface {\n", indent, ig.name, typeParamStmt)
	for _, sig := range ig.funcSignatures {
		signatureStr, err := sig.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += fmt.Sprintf("%s\t%s\n", indent, signatureStr)
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
