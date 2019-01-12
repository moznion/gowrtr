package gowrtr

import "fmt"

type ForGenerator struct {
	Condition  string
	Statements []CodeGeneratable
}

func NewForGenerator(condition string, statements ...CodeGeneratable) *ForGenerator {
	return &ForGenerator{
		Condition:  condition,
		Statements: statements,
	}
}

func (fg *ForGenerator) AddStatements(statements ...CodeGeneratable) *ForGenerator {
	return &ForGenerator{
		Condition:  fg.Condition,
		Statements: append(fg.Statements, statements...),
	}
}

func (fg *ForGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	cond := fg.Condition
	stmt := fmt.Sprintf("%sfor %s", indent, cond)
	if cond != "" {
		stmt += " "
	}
	stmt += "{\n"

	nextIndentLevel := indentLevel + 1
	for _, c := range fg.Statements {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
