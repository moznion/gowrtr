package generator

import "fmt"

// DefaultCase represents a code generator for `default` block of `switch-case` notation.
type DefaultCase struct {
	statements []Statement
}

// NewDefaultCase returns a new `DefaultCase`.
func NewDefaultCase(statements ...Statement) *DefaultCase {
	return &DefaultCase{
		statements: statements,
	}
}

// AddStatements adds statements for `default` block to `DefaultCase`. This does *not* set, just add.
// This method returns a *new* `DefaultCase`; it means this method acts as immutable.
func (d *DefaultCase) AddStatements(statements ...Statement) *DefaultCase {
	return &DefaultCase{
		statements: append(d.statements, statements...),
	}
}

// Statements sets statements for `default` block to `DefaultCase`. This does *not* add, just set.
// This method returns a *new* `DefaultCase`; it means this method acts as immutable.
func (d *DefaultCase) Statements(statements ...Statement) *DefaultCase {
	return &DefaultCase{
		statements: statements,
	}
}

// Generate generates `default` block as golang code.
func (d *DefaultCase) Generate(indentLevel int) (string, error) {
	indent := BuildIndent(indentLevel)
	nextIndentLevel := indentLevel + 1

	stmt := fmt.Sprintf("%sdefault:\n", indent)
	for _, statement := range d.statements {
		gen, err := statement.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	return stmt, nil
}
