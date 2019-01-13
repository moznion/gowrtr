package generator

import "fmt"

// PackageGenerator represents a code generator for `package` statement.
type PackageGenerator struct {
	Name string
}

// NewPackageGenerator returns a new `PackageGenerator`.
func NewPackageGenerator(packageName string) *PackageGenerator {
	return &PackageGenerator{
		Name: packageName,
	}
}

// Generate generates a package statement.
func (pg *PackageGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)
	return fmt.Sprintf("%spackage %s\n", indent, pg.Name), nil
}
