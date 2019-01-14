package generator

import "fmt"

// DefaultCase represents a code generator for `default` block of `switch-case` notation.
type DefaultCase struct {
	Statements []Statement
}

// NewDefaultCase returns a new `DefaultCase`.
func NewDefaultCase(statements ...Statement) *DefaultCase {
	return &DefaultCase{
		Statements: statements,
	}
}

// AddStatements adds statements for `default` block to `DefaultCase`.
// This method returns a *new* `DefaultCase`; it means this method acts as immutable.
func (d *DefaultCase) AddStatements(statements ...Statement) *DefaultCase {
	return &DefaultCase{
		Statements: append(d.Statements, statements...),
	}
}

// Generate generates `default` block as golang code.
func (d *DefaultCase) Generate(indentLevel int) (string, error) {
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
