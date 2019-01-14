package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// StructField represents a field of the struct.
type StructField struct {
	Name string
	Type string
	Tag  string
}

// Struct represents a code generator for `struct` notation.
type Struct struct {
	Name   string
	Fields []*StructField
}

// NewStruct returns a new `Struct`.
func NewStruct(name string) *Struct {
	return &Struct{
		Name: name,
	}
}

// AddField adds a struct field to `Struct`.
// This method returns a *new* `Struct`; it means this method acts as immutable.
func (sg *Struct) AddField(name string, typ string, tag ...string) *Struct {
	l := len(tag)
	t := ""
	if l > 0 {
		t = tag[0]
	}

	return &Struct{
		Name: sg.Name,
		Fields: append(sg.Fields, &StructField{
			Name: name,
			Type: typ,
			Tag:  t,
		}),
	}
}

// Generate generates `struct` block as golang code.
func (sg *Struct) Generate(indentLevel int) (string, error) {
	if sg.Name == "" {
		return "", errmsg.StructNameIsNilErr()
	}

	indent := buildIndent(indentLevel)
	stmt := fmt.Sprintf("%stype %s struct {\n", indent, sg.Name)

	for _, field := range sg.Fields {
		if field.Name == "" {
			return "", errmsg.StructFieldNameIsEmptyErr()
		}
		if field.Type == "" {
			return "", errmsg.StructFieldTypeIsEmptyErr()
		}

		stmt += fmt.Sprintf("%s\t%s %s", indent, field.Name, field.Type)
		if tag := field.Tag; tag != "" {
			stmt += fmt.Sprintf(" `%s`", tag)
		}
		stmt += "\n"
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
