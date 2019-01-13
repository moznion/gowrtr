package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// CaseGenerator represents a generator of `case` statement.
// See also: https://tour.golang.org/flowcontrol/9
type CaseGenerator struct {
	Condition  string
	Statements []StatementGenerator
}

// NewCaseGenerator creates a new `CaseGenerator`.
func NewCaseGenerator(condition string, statements ...StatementGenerator) *CaseGenerator {
	return &CaseGenerator{
		Condition:  condition,
		Statements: statements,
	}
}

// AddStatements adds statements to `CaseGenerator`.
// This method returns a *new* CaseGenerator; it means this method acts as immutable.
func (c *CaseGenerator) AddStatements(statements ...StatementGenerator) *CaseGenerator {
	return &CaseGenerator{
		Condition:  c.Condition,
		Statements: append(c.Statements, statements...),
	}
}

// Generate generates `case` statement as golang's code.
func (c *CaseGenerator) Generate(indentLevel int) (string, error) {
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
