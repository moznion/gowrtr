package gowrtr

import "strings"

type ReturnStatementGenerator struct {
	ReturnItems []string
}

func NewReturnStatementGenerator(returnItems ...string) *ReturnStatementGenerator {
	return &ReturnStatementGenerator{
		ReturnItems: returnItems,
	}
}

func (r *ReturnStatementGenerator) AddReturnItems(returnItems ...string) *ReturnStatementGenerator {
	return &ReturnStatementGenerator{
		ReturnItems: append(r.ReturnItems, returnItems...),
	}
}

func (r *ReturnStatementGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := indent + "return"

	if ret := strings.Join(r.ReturnItems, ", "); ret != "" {
		stmt += " " + ret
	}

	return stmt, nil
}
