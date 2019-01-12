package gowrtr

import (
	"fmt"

	"github.com/moznion/gowrtr/errmsg"
)

type StructField struct {
	Name string
	Type string
	Tag  string
}

type StructGenerator struct {
	Name   string
	Fields []*StructField
}

func NewStructGenerator(name string) *StructGenerator {
	return &StructGenerator{
		Name: name,
	}
}

func (sg *StructGenerator) AddField(name string, typ string, tag ...string) *StructGenerator {
	l := len(tag)
	t := ""
	if l > 0 {
		t = tag[0]
	}

	sg.Fields = append(sg.Fields, &StructField{
		Name: name,
		Type: typ,
		Tag:  t,
	})
	return sg
}

func (sg *StructGenerator) GenerateCode(indentLevel int) (string, error) {
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
