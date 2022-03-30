package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// StructField represents a field of the struct.
type StructField struct {
	name string
	typ  string
	tag  string
}

// Struct represents a code generator for `struct` notation.
type Struct struct {
	name           string
	fields         []*StructField
	nameCaller     string
	fieldsCallers  []string
	typeParameters TypeParameters
}

// NewStruct returns a new `Struct`.
func NewStruct(name string) *Struct {
	return &Struct{
		name:       name,
		nameCaller: fetchClientCallerLine(),
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
		name: sg.name,
		fields: append(sg.fields, &StructField{
			name: name,
			typ:  typ,
			tag:  t,
		}),
		nameCaller:     sg.nameCaller,
		fieldsCallers:  append(sg.fieldsCallers, fetchClientCallerLine()),
		typeParameters: sg.typeParameters,
	}
}

// TypeParameters sets the TypeParameters onto the current caller Struct.
func (sg *Struct) TypeParameters(typeParameters TypeParameters) *Struct {
	return &Struct{
		name:           sg.name,
		fields:         sg.fields,
		nameCaller:     sg.nameCaller,
		fieldsCallers:  sg.fieldsCallers,
		typeParameters: typeParameters,
	}
}

// Generate generates `struct` block as golang code.
func (sg *Struct) Generate(indentLevel int) (string, error) {
	if sg.name == "" {
		return "", errmsg.StructNameIsNilErr(sg.nameCaller)
	}

	indent := BuildIndent(indentLevel)

	typeParametersStmt := ""
	if sg.typeParameters != nil && len(sg.typeParameters) > 0 {
		var err error
		typeParametersStmt, err = sg.typeParameters.Generate(indentLevel)
		if err != nil {
			return "", err
		}
	}

	stmt := fmt.Sprintf("%stype %s%s struct {\n", indent, sg.name, typeParametersStmt)

	for i, field := range sg.fields {
		if field.name == "" {
			return "", errmsg.StructFieldNameIsEmptyErr(sg.fieldsCallers[i])
		}
		if field.typ == "" {
			return "", errmsg.StructFieldTypeIsEmptyErr(sg.fieldsCallers[i])
		}

		stmt += fmt.Sprintf("%s\t%s %s", indent, field.name, field.typ)
		if tag := field.tag; tag != "" {
			stmt += fmt.Sprintf(" `%s`", tag)
		}
		stmt += "\n"
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
