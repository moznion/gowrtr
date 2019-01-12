package gowrtr

import (
	"fmt"
)

type ImportGenerator struct {
	Names []string
}

func NewImportGenerator(names ...string) *ImportGenerator {
	return &ImportGenerator{
		Names: names,
	}
}

func (ig *ImportGenerator) AddImport(imp string) *ImportGenerator {
	return &ImportGenerator{
		Names: append(ig.Names, imp),
	}
}

func (ig *ImportGenerator) Generate(indentLevel int) (string, error) {
	if len(ig.Names) <= 0 {
		return "", nil
	}

	indent := buildIndent(indentLevel)
	stmt := indent + "import (\n"
	for _, name := range ig.Names {
		stmt += fmt.Sprintf("%s\t\"%s\"\n", indent, name)
	}
	stmt += indent + ")\n"

	return stmt, nil
}
