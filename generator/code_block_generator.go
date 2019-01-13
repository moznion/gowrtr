package gowrtr

type CodeBlockGenerator struct {
	Statements []StatementGenerator
}

func NewCodeBlockGenerator(statements ...StatementGenerator) *CodeBlockGenerator {
	return &CodeBlockGenerator{
		Statements: statements,
	}
}

func (c *CodeBlockGenerator) AddStatements(statements ...StatementGenerator) *CodeBlockGenerator {
	return &CodeBlockGenerator{
		Statements: append(c.Statements, statements...),
	}
}

func (c *CodeBlockGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := indent + "{\n"

	nextIndentLevel := indentLevel + 1
	for _, generator := range c.Statements {
		gen, err := generator.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += indent + "}\n"
	return stmt, nil
}
