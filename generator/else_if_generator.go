package generator

import "fmt"

// ElseIfGenerator represents a code generator for `else-if` block.
type ElseIfGenerator struct {
	Condition  string
	Statements []StatementGenerator
}

// NewElseIfGenerator returns a new `ElseIfGenerator`.
func NewElseIfGenerator(condition string, statements ...StatementGenerator) *ElseIfGenerator {
	return &ElseIfGenerator{
		Condition:  condition,
		Statements: statements,
	}
}

// AddStatements adds statements for the `else-if` block to `ElseIfGenerator`.
// This method returns a *new* `ElseIfGenerator`; it means this method acts as immutable.
func (ig *ElseIfGenerator) AddStatements(statements ...StatementGenerator) *ElseIfGenerator {
	return &ElseIfGenerator{
		Condition:  ig.Condition,
		Statements: append(ig.Statements, statements...),
	}
}

// Generate generates `else-if` block as golang's code.
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
