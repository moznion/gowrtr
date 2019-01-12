package gowrtr

import (
	"fmt"
)

type IfGenerator struct {
	Condition  string
	Statements []CodeGenerator
}

func NewIfGenerator(condition string, statements ...CodeGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:  condition,
		Statements: statements,
	}
}

func (ig *IfGenerator) AddStatements(statements ...CodeGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:  ig.Condition,
		Statements: append(ig.Statements, statements...),
	}
}

func (ig *IfGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := fmt.Sprintf("%sif %s {\n", indent, ig.Condition)

	nextIndentLevel := indentLevel + 1
	for _, c := range ig.Statements {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
