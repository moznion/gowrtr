package generator

import "fmt"

// SwitchGenerator represents a code generator for `switch` statement.
// See also: https://tour.golang.org/flowcontrol/9
type SwitchGenerator struct {
	Condition        string
	CaseStatements   []*CaseGenerator
	DefaultStatement *DefaultCaseGenerator
}

// NewSwitchGenerator returns a new `SwitchGenerator`.
func NewSwitchGenerator(condition string) *SwitchGenerator {
	return &SwitchGenerator{
		Condition: condition,
	}
}

// AddCaseStatements adds `case` statements to `SwitchGenerator`.
// This method returns a *new* `SwitchGenerator`; it means this method acts as immutable.
func (s *SwitchGenerator) AddCaseStatements(statements ...*CaseGenerator) *SwitchGenerator {
	return &SwitchGenerator{
		Condition:        s.Condition,
		CaseStatements:   append(s.CaseStatements, statements...),
		DefaultStatement: s.DefaultStatement,
	}
}

// SetDefaultStatement sets a `default` statement to `SwitchGenerator`.
// This method returns a *new* `SwitchGenerator`; it means this method acts as immutable.
func (s *SwitchGenerator) SetDefaultStatement(statement *DefaultCaseGenerator) *SwitchGenerator {
	return &SwitchGenerator{
		Condition:        s.Condition,
		CaseStatements:   s.CaseStatements,
		DefaultStatement: statement,
	}
}

// Generate generates `switch` statement as golang's code.
func (s *SwitchGenerator) Generate(indentLevel int) (string, error) {
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
