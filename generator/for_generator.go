package generator

import "fmt"

// ForGenerator represents a code generator for `for` block.
type ForGenerator struct {
	Condition  string
	Statements []StatementGenerator
}

// NewForGenerator returns a new `ForGenerator`.
func NewForGenerator(condition string, statements ...StatementGenerator) *ForGenerator {
	return &ForGenerator{
		Condition:  condition,
		Statements: statements,
	}
}

// AddStatements adds statements for `for` block to `ForGenerator`.
// This method returns a *new* `ForGenerator`; it means this method acts as immutable.
func (fg *ForGenerator) AddStatements(statements ...StatementGenerator) *ForGenerator {
	return &ForGenerator{
		Condition:  fg.Condition,
		Statements: append(fg.Statements, statements...),
	}
}

// Generate generates a `for` block as golang code.
func (fg *ForGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	cond := fg.Condition
	stmt := fmt.Sprintf("%sfor %s", indent, cond)
	if cond != "" {
		stmt += " "
	}
	stmt += "{\n"

	nextIndentLevel := indentLevel + 1
	for _, c := range fg.Statements {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
