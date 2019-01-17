package generator

import "fmt"

// Switch represents a code generator for `switch` statement.
// See also: https://tour.golang.org/flowcontrol/9
type Switch struct {
	condition        string
	caseStatements   []*Case
	defaultStatement *DefaultCase
}

// NewSwitch returns a new `Switch`.
func NewSwitch(condition string) *Switch {
	return &Switch{
		condition: condition,
	}
}

// AddCase adds `case` statements to `Switch`. This does *not* set, just add.
// This method returns a *new* `Switch`; it means this method acts as immutable.
func (s *Switch) AddCase(statements ...*Case) *Switch {
	return &Switch{
		condition:        s.condition,
		caseStatements:   append(s.caseStatements, statements...),
		defaultStatement: s.defaultStatement,
	}
}

// Case sets `case` statements to `Switch`. This does *not* add, just set.
// This method returns a *new* `Switch`; it means this method acts as immutable.
func (s *Switch) Case(statements ...*Case) *Switch {
	return &Switch{
		condition:        s.condition,
		caseStatements:   statements,
		defaultStatement: s.defaultStatement,
	}
}

// Default sets a `default` statement to `Switch`.
// This method returns a *new* `Switch`; it means this method acts as immutable.
func (s *Switch) Default(statement *DefaultCase) *Switch {
	return &Switch{
		condition:        s.condition,
		caseStatements:   s.caseStatements,
		defaultStatement: statement,
	}
}

// Generate generates `switch` statement as golang code.
func (s *Switch) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := fmt.Sprintf("%sswitch %s {\n", indent, s.condition)
	for _, statement := range s.caseStatements {
		if statement == nil {
			continue
		}
		gen, err := statement.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	if defaultStatement := s.defaultStatement; defaultStatement != nil {
		gen, err := defaultStatement.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
