package generator

import "fmt"

type DefaultCaseGenerator struct {
	Statements []StatementGenerator
}

func NewDefaultCaseGenerator(statements ...StatementGenerator) *DefaultCaseGenerator {
	return &DefaultCaseGenerator{
		Statements: statements,
	}
}

func (d *DefaultCaseGenerator) AddStatements(statements ...StatementGenerator) *DefaultCaseGenerator {
	return &DefaultCaseGenerator{
		Statements: append(d.Statements, statements...),
	}
}

func (d *DefaultCaseGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)
	nextIndentLevel := indentLevel + 1

	stmt := fmt.Sprintf("%sdefault:\n", indent)
	for _, statement := range d.Statements {
		gen, err := statement.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	return stmt, nil
}
