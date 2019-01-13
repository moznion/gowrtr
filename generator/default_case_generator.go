package generator

import "fmt"

// DefaultCaseGenerator represents a code generator for `default` block of `switch-case` notation.
type DefaultCaseGenerator struct {
	Statements []StatementGenerator
}

// NewDefaultCaseGenerator returns a new `DefaultCaseGenerator`.
func NewDefaultCaseGenerator(statements ...StatementGenerator) *DefaultCaseGenerator {
	return &DefaultCaseGenerator{
		Statements: statements,
	}
}

// AddStatements adds statements for `default` block to `DefaultCaseGenerator`.
// This method returns a *new* `DefaultCaseGenerator`; it means this method acts as immutable.
func (d *DefaultCaseGenerator) AddStatements(statements ...StatementGenerator) *DefaultCaseGenerator {
	return &DefaultCaseGenerator{
		Statements: append(d.Statements, statements...),
	}
}

// Generate generates `default` block as golang code.
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
