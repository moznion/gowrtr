package generator

import (
	"fmt"
)

// Import represents a code generator for `import` statement.
type Import struct {
	names []string
}

// NewImport returns a new `Import`.
func NewImport(names ...string) *Import {
	return &Import{
		names: names,
	}
}

// AddImports adds import items to `Import`. This does *not* set, just add.
// This method returns a *new* `Import`; it means this method acts as immutable.
func (ig *Import) AddImports(imps ...string) *Import {
	return &Import{
		names: append(ig.names, imps...),
	}
}

// Imports sets import items to `Import`. This does *not* add, just set.
// This method returns a *new* `Import`; it means this method acts as immutable.
func (ig *Import) Imports(imps ...string) *Import {
	return &Import{
		names: imps,
	}
}

// Generate generates `import` statement as golang code.
func (ig *Import) Generate(indentLevel int) (string, error) {
	if len(ig.names) <= 0 {
		return "", nil
	}

	indent := BuildIndent(indentLevel)
	stmt := indent + "import (\n"
	for _, name := range ig.names {
		if name == "" {
			continue
		}
		stmt += fmt.Sprintf("%s\t\"%s\"\n", indent, name)
	}
	stmt += indent + ")\n"

	return stmt, nil
}
