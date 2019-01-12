package gowrtr

import "fmt"

type PackageGenerator struct {
	Name       string
	generators []CodeGenerator
}

func NewPackageGenerator(packageName string) *PackageGenerator {
	return &PackageGenerator{
		Name: packageName,
	}
}

func (pg *PackageGenerator) AddStatement(c CodeGenerator) *PackageGenerator {
	pg.generators = append(pg.generators, c)
	return pg
}

func (pg *PackageGenerator) Generate(indentLevel int) (string, error) {
	return fmt.Sprintf("package %s\n", pg.Name), nil
}
