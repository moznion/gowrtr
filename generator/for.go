package generator

import "fmt"

// For represents a code generator for `for` block.
type For struct {
	condition  string
	statements []Statement
}

// NewFor returns a new `For`.
func NewFor(condition string, statements ...Statement) *For {
	return &For{
		condition:  condition,
		statements: statements,
	}
}

// AddStatements adds statements for `for` block to `For`. This does *not* set, just add.
// This method returns a *new* `For`; it means this method acts as immutable.
func (fg *For) AddStatements(statements ...Statement) *For {
	return &For{
		condition:  fg.condition,
		statements: append(fg.statements, statements...),
	}
}

// Statements sets statements for `for` block to `For`. This does *not* add, just set.
// This method returns a *new* `For`; it means this method acts as immutable.
func (fg *For) Statements(statements ...Statement) *For {
	return &For{
		condition:  fg.condition,
		statements: statements,
	}
}

// Generate generates a `for` block as golang code.
func (fg *For) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	cond := fg.condition
	stmt := fmt.Sprintf("%sfor %s", indent, cond)
	if cond != "" {
		stmt += " "
	}
	stmt += "{\n"

	nextIndentLevel := indentLevel + 1
	for _, c := range fg.statements {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
