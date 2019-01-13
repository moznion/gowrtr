package generator

import (
	"fmt"
)

// ImportGenerator represents a code generator for `import` statement.
type ImportGenerator struct {
	Names []string
}

// NewImportGenerator returns a new `ImportGenerator`.
func NewImportGenerator(names ...string) *ImportGenerator {
	return &ImportGenerator{
		Names: names,
	}
}

// AddImports adds import items to `ImportGenerator`.
// This method returns a *new* `ImportGenerator`; it means this method acts as immutable.
func (ig *ImportGenerator) AddImports(imps ...string) *ImportGenerator {
	return &ImportGenerator{
		Names: append(ig.Names, imps...),
	}
}

// Generate generates `import` statement as golang code.
func (ig *ImportGenerator) Generate(indentLevel int) (string, error) {
	if len(ig.Names) <= 0 {
		return "", nil
	}

	indent := buildIndent(indentLevel)
	stmt := indent + "import (\n"
	for _, name := range ig.Names {
		if name == "" {
			continue
		}
		stmt += fmt.Sprintf("%s\t\"%s\"\n", indent, name)
	}
	stmt += indent + ")\n"

	return stmt, nil
}
