package generator

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// TypeParameter is a "generic" type parameter. e,g. `T string`; T is a parameter and int is type.
type TypeParameter struct {
	parameter string
	types     []string
	caller    string
}

// NewTypeParameter returns a new TypeParameter.
func NewTypeParameter(parameter string, typ string) *TypeParameter {
	return NewTypeParameters(parameter, []string{typ})
}

// NewTypeParameters returns a new TypeParameter. If the `types` argument has the multiple types, those types become a union type.
func NewTypeParameters(parameter string, types []string) *TypeParameter {
	return &TypeParameter{
		parameter: parameter,
		types:     types,
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

		if len(tp.types) <= 0 {
			return "", errmsg.TypeParameterTypeIsEmptyErr(tp.caller)
		}

		isBlank := true
		for _, typ := range tp.types {
			if typ != "" {
				isBlank = false
				break
			}
		}
		if isBlank {
			return "", errmsg.TypeParameterTypeIsEmptyErr(tp.caller)
		}

		typeParameterStmts[i] = tp.parameter + " " + strings.Join(tp.types, " | ")
	}

	return "[" + strings.Join(typeParameterStmts, ", ") + "]", nil
}

// TypeArguments is an array of type argument for go generics.
type TypeArguments []string

// Generate generates the type-arguments as golang code.
func (tas TypeArguments) Generate(indentLevel int) (string, error) {
	return "[" + strings.Join(tas, ", ") + "]", nil
}
