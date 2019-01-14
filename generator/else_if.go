package generator

import "fmt"

// ElseIf represents a code generator for `else-if` block.
type ElseIf struct {
	Condition  string
	Statements []Statement
}

// NewElseIf returns a new `ElseIf`.
func NewElseIf(condition string, statements ...Statement) *ElseIf {
	return &ElseIf{
		Condition:  condition,
		Statements: statements,
	}
}

// AddStatements adds statements for the `else-if` block to `ElseIf`.
// This method returns a *new* `ElseIf`; it means this method acts as immutable.
func (ig *ElseIf) AddStatements(statements ...Statement) *ElseIf {
	return &ElseIf{
		Condition:  ig.Condition,
		Statements: append(ig.Statements, statements...),
	}
}

// Generate generates `else-if` block as golang code.
func (ig *ElseIf) Generate(indentLevel int) (string, error) {
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
