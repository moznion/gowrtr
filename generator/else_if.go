package generator

import "fmt"

// ElseIf represents a code generator for `else-if` block.
type ElseIf struct {
	condition  string
	statements []Statement
}

// NewElseIf returns a new `ElseIf`.
func NewElseIf(condition string, statements ...Statement) *ElseIf {
	return &ElseIf{
		condition:  condition,
		statements: statements,
	}
}

// AddStatements adds statements for the `else-if` block to `ElseIf`. This does *not* set, just add.
// This method returns a *new* `ElseIf`; it means this method acts as immutable.
func (ei *ElseIf) AddStatements(statements ...Statement) *ElseIf {
	return &ElseIf{
		condition:  ei.condition,
		statements: append(ei.statements, statements...),
	}
}

// Statements sets statements for the `else-if` block to `ElseIf`. This does *not* add, just set.
// This method returns a *new* `ElseIf`; it means this method acts as immutable.
func (ei *ElseIf) Statements(statements ...Statement) *ElseIf {
	return &ElseIf{
		condition:  ei.condition,
		statements: statements,
	}
}

// Generate generates `else-if` block as golang code.
func (ei *ElseIf) Generate(indentLevel int) (string, error) {
	indent := BuildIndent(indentLevel)

	stmt := fmt.Sprintf(" else if %s {\n", ei.condition)

	nextIndentLevel := indentLevel + 1
	for _, c := range ei.statements {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}", indent)

	return stmt, nil
}
