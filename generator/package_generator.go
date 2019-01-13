package generator

import "fmt"

type PackageGenerator struct {
	Name string
}

func NewPackageGenerator(packageName string) *PackageGenerator {
	return &PackageGenerator{
		Name: packageName,
	}
}

func (pg *PackageGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)
	return fmt.Sprintf("%spackage %s\n", indent, pg.Name), nil
}
