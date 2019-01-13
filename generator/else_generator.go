package generator

import "fmt"

// ElseGenerator represents a code generator for `else` block.
type ElseGenerator struct {
	Statements []StatementGenerator
}

// NewElseGenerator returns a new `ElseGenerator`.
func NewElseGenerator(statements ...StatementGenerator) *ElseGenerator {
	return &ElseGenerator{
		Statements: statements,
	}
}

// AddStatements adds statements for `else` block to `ElseGenerator`.
// This method returns a *new* `ElseGenerator`; it means this method acts as immutable.
func (ig *ElseGenerator) AddStatements(statements ...StatementGenerator) *ElseGenerator {
	return &ElseGenerator{
		Statements: append(ig.Statements, statements...),
	}
}

// Generate generates `else` block as golang's code.
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
