package generator

import "strings"

// ReturnStatement represents a code generator for `return` statement.
type ReturnStatement struct {
	returnItems []string
}

// NewReturnStatement returns a new `ReturnStatement`.
func NewReturnStatement(returnItems ...string) *ReturnStatement {
	return &ReturnStatement{
		returnItems: returnItems,
	}
}

// AddReturnItems adds return items to `ReturnStatement`. This does *not* set, just add.
// This method returns a *new* `ReturnStatement`; it means this method acts as immutable.
func (r *ReturnStatement) AddReturnItems(returnItems ...string) *ReturnStatement {
	return &ReturnStatement{
		returnItems: append(r.returnItems, returnItems...),
	}
}

// ReturnItems sets return items to `ReturnStatement`. This does *not* add, just set.
// This method returns a *new* `ReturnStatement`; it means this method acts as immutable.
func (r *ReturnStatement) ReturnItems(returnItems ...string) *ReturnStatement {
	return &ReturnStatement{
		returnItems: returnItems,
	}
}

// Generate generates `return` statement as golang code.
func (r *ReturnStatement) Generate(indentLevel int) (string, error) {
	indent := BuildIndent(indentLevel)

	stmt := indent + "return"
	if ret := strings.Join(r.returnItems, ", "); ret != "" {
		stmt += " " + ret
	}
	stmt += "\n"

	return stmt, nil
}
