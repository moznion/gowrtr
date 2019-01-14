package generator

import "fmt"

// Else represents a code generator for `else` block.
type Else struct {
	Statements []Statement
}

// NewElse returns a new `Else`.
func NewElse(statements ...Statement) *Else {
	return &Else{
		Statements: statements,
	}
}

// AddStatements adds statements for `else` block to `Else`.
// This method returns a *new* `Else`; it means this method acts as immutable.
func (ig *Else) AddStatements(statements ...Statement) *Else {
	return &Else{
		Statements: append(ig.Statements, statements...),
	}
}

// Generate generates `else` block as golang code.
func (ig *Else) Generate(indentLevel int) (string, error) {
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
