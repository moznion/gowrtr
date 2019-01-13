package generator

import "fmt"

type SwitchGenerator struct {
	Condition        string
	CaseStatements   []*CaseGenerator
	DefaultStatement *DefaultCaseGenerator
}

func NewSwitchGenerator(condition string) *SwitchGenerator {
	return &SwitchGenerator{
		Condition: condition,
	}
}

func (s *SwitchGenerator) AddCaseStatements(statements ...*CaseGenerator) *SwitchGenerator {
	return &SwitchGenerator{
		Condition:        s.Condition,
		CaseStatements:   append(s.CaseStatements, statements...),
		DefaultStatement: s.DefaultStatement,
	}
}

func (s *SwitchGenerator) SetDefaultStatement(statement *DefaultCaseGenerator) *SwitchGenerator {
	return &SwitchGenerator{
		Condition:        s.Condition,
		CaseStatements:   s.CaseStatements,
		DefaultStatement: statement,
	}
}

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
