package gowrtr

type CodeBlockGenerator struct {
	Generators []CodeGenerator
}

func NewCodeBlockGenerator(generators ...CodeGenerator) *CodeBlockGenerator {
	return &CodeBlockGenerator{
		Generators: generators,
	}
}

func (c *CodeBlockGenerator) AddStatements(generators ...CodeGenerator) *CodeBlockGenerator {
	c.Generators = append(c.Generators, generators...)
	return c
}

func (c *CodeBlockGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := indent + "{\n"

	nextIndentLevel := indentLevel + 1
	for _, generator := range c.Generators {
		gen, err := generator.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += indent + "}\n"
	return stmt, nil
}
