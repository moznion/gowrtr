package generator

import (
	"fmt"
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

type compositeLiteralField struct {
	key   string
	value Statement
}

// CompositeLiteral represents a code generator for composite literal.
// Please see also: https://golang.org/doc/effective_go.html#composite_literals
type CompositeLiteral struct {
	typ    string
	fields []*compositeLiteralField
}

// NewCompositeLiteral returns a new `CompositeLiteral`.
func NewCompositeLiteral(typ string) *CompositeLiteral {
	return &CompositeLiteral{
		typ: typ,
	}
}

// AddField adds a field as `Statement` to `ComposeLiteral`.
// This method returns a *new* `Struct`; it means this method acts as immutable.
func (c *CompositeLiteral) AddField(key string, value Statement) *CompositeLiteral {
	return &CompositeLiteral{
		typ: c.typ,
		fields: append(c.fields, &compositeLiteralField{
			key:   key,
			value: value,
		}),
	}
}

// AddFieldStr adds a field as string to `ComposeLiteral`.
// This method returns a *new* `Struct`; it means this method acts as immutable.
func (c *CompositeLiteral) AddFieldStr(key string, value string) *CompositeLiteral {
	return &CompositeLiteral{
		typ: c.typ,
		fields: append(c.fields, &compositeLiteralField{
			key:   key,
			value: NewRawStatement(fmt.Sprintf(`"%s"`, value)),
		}),
	}
}

// AddFieldRaw adds a field as raw text to `ComposeLiteral`.
// This method returns a *new* `Struct`; it means this method acts as immutable.
func (c *CompositeLiteral) AddFieldRaw(key string, value interface{}) *CompositeLiteral {
	return &CompositeLiteral{
		typ: c.typ,
		fields: append(c.fields, &compositeLiteralField{
			key:   key,
			value: NewRawStatement(fmt.Sprintf("%v", value)),
		}),
	}
}

// Generate generates composite literal block as golang code.
func (c *CompositeLiteral) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)
	nextLevelIndent := buildIndent(indentLevel + 1)

	stmt := fmt.Sprintf("%s%s{\n", indent, c.typ)
	for _, field := range c.fields {
		genValue, err := field.value.Generate(indentLevel + 1)
		if err != nil {
			return "", err
		}

		genValue = strings.TrimSpace(genValue)

		stmt += nextLevelIndent

		if key := field.key; key != "" {
			stmt += key + ": "
		}
		if genValue == "" {
			return "", errmsg.ValueOfCompositeLiteralIsEmptyError()
		}
		stmt += fmt.Sprintf("%s,\n", genValue)
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
