package generator

import (
	"fmt"
)

// Import represents a code generator for `import` statement.
type Import struct {
	Names []string
}

// NewImport returns a new `Import`.
func NewImport(names ...string) *Import {
	return &Import{
		Names: names,
	}
}

// AddImports adds import items to `Import`.
// This method returns a *new* `Import`; it means this method acts as immutable.
func (ig *Import) AddImports(imps ...string) *Import {
	return &Import{
		Names: append(ig.Names, imps...),
	}
}

// Generate generates `import` statement as golang code.
func (ig *Import) Generate(indentLevel int) (string, error) {
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
