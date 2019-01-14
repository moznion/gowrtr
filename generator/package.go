package generator

import "fmt"

// Package represents a code generator for `package` statement.
type Package struct {
	Name string
}

// NewPackage returns a new `Package`.
func NewPackage(packageName string) *Package {
	return &Package{
		Name: packageName,
	}
}

// Generate generates a package statement.
func (pg *Package) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)
	return fmt.Sprintf("%spackage %s\n", indent, pg.Name), nil
}
