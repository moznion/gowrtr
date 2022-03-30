package generator

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// TypeParameter is a "generic" type parameter. e,g. `T string`; T is a parameter and int is type.
type TypeParameter struct {
	parameter string
	typ       string
	caller    string
}

// NewTypeParameter returns a new TypeParameter.
func NewTypeParameter(parameter string, typ string) *TypeParameter {
	return &TypeParameter{
		parameter: parameter,
		typ:       typ,
		caller:    fetchClientCallerLine(),
	}
}

// TypeParameters is an array of TypeParameter pointers for go generics.
type TypeParameters []*TypeParameter

// Generate generates the type-parameters as golang code.
func (tps TypeParameters) Generate(indentLevel int) (string, error) {
	typeParameterStmts := make([]string, len(tps))
	for i, tp := range tps {
		if tp.parameter == "" {
			return "", errmsg.TypeParameterParameterIsEmptyErr(tp.caller)
		}

		if tp.typ == "" {
			return "", errmsg.TypeParameterTypeIsEmptyErr(tp.caller)
		}

		typeParameterStmts[i] = tp.parameter + " " + tp.typ
	}

	return "[" + strings.Join(typeParameterStmts, ", ") + "]", nil
}

// TypeArguments is an array of type argument for go generics.
type TypeArguments []string

// Generate generates the type-arguments as golang code.
func (tas TypeArguments) Generate(indentLevel int) (string, error) {
	return "[" + strings.Join(tas, ", ") + "]", nil
}
