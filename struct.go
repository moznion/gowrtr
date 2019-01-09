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

type Struct struct {
	Name   string
	Fields []*StructField
}

func NewStruct(name string, fields []*StructField) (*Struct, error) {
	if name == "" {
		return nil, errmsg.StructNameIsNilErr()
	}

	for _, field := range fields {
		if field.Name == "" {
			return nil, errmsg.StructFieldNameIsEmptyErr()
		}
		if field.Type == "" {
			return nil, errmsg.StructFieldTypeIsEmptyErr()
		}
	}

	return &Struct{
		Name:   name,
		Fields: fields,
	}, nil
}

func (s *Struct) String() string {
	stmt := fmt.Sprintf("type %s struct {\n", s.Name)
	for _, field := range s.Fields {
		stmt += fmt.Sprintf("\t%s %s", field.Name, field.Type)
		if tag := field.Tag; tag != "" {
			stmt += fmt.Sprintf(" `%s`", tag)
		}
		stmt += "\n"
	}
	stmt += "}"

	return stmt
}
