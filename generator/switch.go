package generator

import "fmt"

// Switch represents a code generator for `switch` statement.
// See also: https://tour.golang.org/flowcontrol/9
type Switch struct {
	Condition        string
	CaseStatements   []*Case
	DefaultStatement *DefaultCase
}

// NewSwitch returns a new `Switch`.
func NewSwitch(condition string) *Switch {
	return &Switch{
		Condition: condition,
	}
}

// AddCaseStatements adds `case` statements to `Switch`.
// This method returns a *new* `Switch`; it means this method acts as immutable.
func (s *Switch) AddCaseStatements(statements ...*Case) *Switch {
	return &Switch{
		Condition:        s.Condition,
		CaseStatements:   append(s.CaseStatements, statements...),
		DefaultStatement: s.DefaultStatement,
	}
}

// SetDefaultStatement sets a `default` statement to `Switch`.
// This method returns a *new* `Switch`; it means this method acts as immutable.
func (s *Switch) SetDefaultStatement(statement *DefaultCase) *Switch {
	return &Switch{
		Condition:        s.Condition,
		CaseStatements:   s.CaseStatements,
		DefaultStatement: statement,
	}
}

// Generate generates `switch` statement as golang code.
func (s *Switch) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := fmt.Sprintf("%sswitch %s {\n", indent, s.Condition)
	for _, statement := range s.CaseStatements {
		if statement == nil {
			continue
		}
		gen, err := statement.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	if defaultStatement := s.DefaultStatement; defaultStatement != nil {
		gen, err := defaultStatement.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
