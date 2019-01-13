package generator

import "fmt"

type ElseIfGenerator struct {
	Condition  string
	Statements []StatementGenerator
}

func NewElseIfGenerator(condition string, statements ...StatementGenerator) *ElseIfGenerator {
	return &ElseIfGenerator{
		Condition:  condition,
		Statements: statements,
	}
}

func (ig *ElseIfGenerator) AddStatements(statements ...StatementGenerator) *ElseIfGenerator {
	return &ElseIfGenerator{
		Condition:  ig.Condition,
		Statements: append(ig.Statements, statements...),
	}
}

func (ig *ElseIfGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := fmt.Sprintf(" else if %s {\n", ig.Condition)

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
