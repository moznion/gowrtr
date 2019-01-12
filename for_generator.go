package gowrtr

import "fmt"

type ForGenerator struct {
	Condition  string
	Generators []CodeGenerator
}

func NewForGenerator(condition string, generators ...CodeGenerator) *ForGenerator {
	return &ForGenerator{
		Condition:  condition,
		Generators: generators,
	}
}

func (fg *ForGenerator) AddStatements(c ...CodeGenerator) *ForGenerator {
	fg.Generators = append(fg.Generators, c...)
	return fg
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
	for _, c := range fg.Generators {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}
	stmt += fmt.Sprintf("%s}\n", indent)

	return stmt, nil
}
