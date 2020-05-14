package generator

import "fmt"

// Else represents a code generator for `else` block.
type Else struct {
	statements []Statement
}

// NewElse returns a new `Else`.
func NewElse(statements ...Statement) *Else {
	return &Else{
		statements: statements,
	}
}

// AddStatements adds statements for `else` block to `Else`. This does *not* set, just add.
// This method returns a *new* `Else`; it means this method acts as immutable.
func (e *Else) AddStatements(statements ...Statement) *Else {
	return &Else{
		statements: append(e.statements, statements...),
	}
}

// Statements sets statements for `else` block to `Else`. This does *not* add, just set.
// This method returns a *new* `Else`; it means this method acts as immutable.
func (e *Else) Statements(statements ...Statement) *Else {
	return &Else{
		statements: statements,
	}
}

// Generate generates `else` block as golang code.
func (e *Else) Generate(indentLevel int) (string, error) {
	stmt := fmt.Sprintf(" else {\n")

	indent := BuildIndent(indentLevel)
	nextIndentLevel := indentLevel + 1
	for _, c := range e.statements {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}", indent)

	return stmt, nil
}
