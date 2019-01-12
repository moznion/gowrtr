package gowrtr

import (
	"fmt"
)

type IfGenerator struct {
	Condition  string
	Generators []CodeGenerator
}

func NewIfGenerator(condition string, generators ...CodeGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:  condition,
		Generators: generators,
	}
}

func (ig *IfGenerator) AddStatements(generators ...CodeGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:  ig.Condition,
		Generators: append(ig.Generators, generators...),
	}
}

func (ig *IfGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := fmt.Sprintf("%sif %s {\n", indent, ig.Condition)

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
