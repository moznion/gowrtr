package generator

import "strings"

// ReturnStatementGenerator represents a code generator for `return` statement.
type ReturnStatementGenerator struct {
	ReturnItems []string
}

// NewReturnStatementGenerator returns a new `ReturnStatementGenerator`.
func NewReturnStatementGenerator(returnItems ...string) *ReturnStatementGenerator {
	return &ReturnStatementGenerator{
		ReturnItems: returnItems,
	}
}

// AddReturnItems adds return items to `ReturnStatement`.
// This method returns a *new* `ReturnStatement`; it means this method acts as immutable.
func (r *ReturnStatementGenerator) AddReturnItems(returnItems ...string) *ReturnStatementGenerator {
	return &ReturnStatementGenerator{
		ReturnItems: append(r.ReturnItems, returnItems...),
	}
}

// Generate generates `return` statement as golang's code.
func (r *ReturnStatementGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := indent + "return"
	if ret := strings.Join(r.ReturnItems, ", "); ret != "" {
		stmt += " " + ret
	}
	stmt += "\n"

	return stmt, nil
}
