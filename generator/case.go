package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// Case represents a code generator for `case` statement.
// See also: https://tour.golang.org/flowcontrol/9
type Case struct {
	Condition  string
	Statements []Statement
}

// NewCase creates a new `Case`.
func NewCase(condition string, statements ...Statement) *Case {
	return &Case{
		Condition:  condition,
		Statements: statements,
	}
}

// AddStatements adds statements to `Case`.
// This method returns a *new* `Case`; it means this method acts as immutable.
func (c *Case) AddStatements(statements ...Statement) *Case {
	return &Case{
		Condition:  c.Condition,
		Statements: append(c.Statements, statements...),
	}
}

// Generate generates `case` statement as golang code.
func (c *Case) Generate(indentLevel int) (string, error) {
	condition := c.Condition
	if condition == "" {
		return "", errmsg.CaseConditionIsEmptyError()
	}

	indent := buildIndent(indentLevel)
	nextIndentLevel := indentLevel + 1

	stmt := fmt.Sprintf("%scase %s:\n", indent, condition)
	for _, statement := range c.Statements {
		gen, err := statement.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	return stmt, nil
}
