package generator

import "fmt"

// For represents a code generator for `for` block.
type For struct {
	Condition  string
	Statements []Statement
}

// NewFor returns a new `For`.
func NewFor(condition string, statements ...Statement) *For {
	return &For{
		Condition:  condition,
		Statements: statements,
	}
}

// AddStatements adds statements for `for` block to `For`.
// This method returns a *new* `For`; it means this method acts as immutable.
func (fg *For) AddStatements(statements ...Statement) *For {
	return &For{
		Condition:  fg.Condition,
		Statements: append(fg.Statements, statements...),
	}
}

// Generate generates a `for` block as golang code.
func (fg *For) Generate(indentLevel int) (string, error) {
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
