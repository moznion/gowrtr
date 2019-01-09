package gowrtr

import (
	"fmt"

	"github.com/moznion/gowrtr/errmsg"
)

// TODO extract
type StructField struct {
	Name string
	Type string
	Tag  string
}

type Struct struct {
	Name   string
	Fields []*StructField
}

func NewStruct(name string, fields []*StructField) *Struct {
	return &Struct{
		Name:   name,
		Fields: fields,
	}
}

func (s *Struct) GenerateCode() (string, error) {
	if s.Name == "" {
		return "", errmsg.StructNameIsNilErr()
	}
	stmt := fmt.Sprintf("type %s struct {\n", s.Name)

	for _, field := range s.Fields {
		if field.Name == "" {
			return "", errmsg.StructFieldNameIsEmptyErr()
		}
		if field.Type == "" {
			return "", errmsg.StructFieldTypeIsEmptyErr()
		}

		stmt += fmt.Sprintf("\t%s %s", field.Name, field.Type)
		if tag := field.Tag; tag != "" {
			stmt += fmt.Sprintf(" `%s`", tag)
		}
		stmt += "\n"
	}
	stmt += "}"

	return stmt, nil
}
