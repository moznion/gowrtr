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
	ig.Names = append(ig.Names, imp)
	return ig
}

func (ig *ImportGenerator) GenerateCode() (string, error) {
	if len(ig.Names) <= 0 {
		return "", nil
	}

	stmt := "import (\n"
	for _, name := range ig.Names {
		stmt += fmt.Sprintf("\t\"%s\"\n", name)
	}
	stmt += ")"

	return stmt, nil
}
