package gowrtr

import (
	"fmt"
)

type IfGenerator struct {
	Cond       string
	Generators []CodeGenerator
}

func NewIfGenerator(cond string, generators ...CodeGenerator) *IfGenerator {
	return &IfGenerator{
		Cond:       cond,
		Generators: generators,
	}
}

func (ig *IfGenerator) AddStatements(generators ...CodeGenerator) *IfGenerator {
	ig.Generators = append(ig.Generators, generators...)
	return ig
}

func (ig *IfGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := fmt.Sprintf("%sif %s {\n", indent, ig.Cond)

	nextIndentLevel := indentLevel + 1
	for _, c := range ig.Generators {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
