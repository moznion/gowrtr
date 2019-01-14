package generator

import "strings"

// ReturnStatement represents a code generator for `return` statement.
type ReturnStatement struct {
	ReturnItems []string
}

// NewReturnStatement returns a new `ReturnStatement`.
func NewReturnStatement(returnItems ...string) *ReturnStatement {
	return &ReturnStatement{
		ReturnItems: returnItems,
	}
}

// AddReturnItems adds return items to `ReturnStatement`.
// This method returns a *new* `ReturnStatement`; it means this method acts as immutable.
func (r *ReturnStatement) AddReturnItems(returnItems ...string) *ReturnStatement {
	return &ReturnStatement{
		ReturnItems: append(r.ReturnItems, returnItems...),
	}
}

// Generate generates `return` statement as golang code.
func (r *ReturnStatement) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := indent + "return"
	if ret := strings.Join(r.ReturnItems, ", "); ret != "" {
		stmt += " " + ret
	}
	stmt += "\n"

	return stmt, nil
}
