package generator

// CodeBlockGenerator represents a plain code block.
//
// example:
// {
//   // do something
//   fmt.Println("blah blah")
// }
type CodeBlockGenerator struct {
	Statements []StatementGenerator
}

// NewCodeBlockGenerator returns a new `CodeBlockGenerator`.
func NewCodeBlockGenerator(statements ...StatementGenerator) *CodeBlockGenerator {
	return &CodeBlockGenerator{
		Statements: statements,
	}
}

// AddStatements adds statements to `CodeBlockGenerator`.
// This method returns a *new* CodeBlockGenerator; it means this method acts as immutable.
func (c *CodeBlockGenerator) AddStatements(statements ...StatementGenerator) *CodeBlockGenerator {
	return &CodeBlockGenerator{
		Statements: append(c.Statements, statements...),
	}
}

// Generate generates plain code block as golang's code.
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
