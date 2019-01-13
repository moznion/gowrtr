package generator

import "fmt"

type ElseGenerator struct {
	Statements []StatementGenerator
}

func NewElseGenerator(statements ...StatementGenerator) *ElseGenerator {
	return &ElseGenerator{
		Statements: statements,
	}
}

func (ig *ElseGenerator) AddStatements(statements ...StatementGenerator) *ElseGenerator {
	return &ElseGenerator{
		Statements: append(ig.Statements, statements...),
	}
}

func (ig *ElseGenerator) Generate(indentLevel int) (string, error) {
	stmt := fmt.Sprintf(" else {\n")

	indent := buildIndent(indentLevel)
	nextIndentLevel := indentLevel + 1
	for _, c := range ig.Statements {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}", indent)

	return stmt, nil
}
