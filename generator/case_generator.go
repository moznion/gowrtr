package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

type CaseGenerator struct {
	Condition  string
	Statements []StatementGenerator
}

func NewCaseGenerator(condition string, statements ...StatementGenerator) *CaseGenerator {
	return &CaseGenerator{
		Condition:  condition,
		Statements: statements,
	}
}

func (c *CaseGenerator) AddStatements(statements ...StatementGenerator) *CaseGenerator {
	return &CaseGenerator{
		Condition:  c.Condition,
		Statements: append(c.Statements, statements...),
	}
}

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
