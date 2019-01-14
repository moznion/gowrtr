package generator

// CodeBlock represents a code generator for plain code block.
//
// example:
// {
//   // do something
//   fmt.Println("blah blah")
// }
type CodeBlock struct {
	Statements []Statement
}

// NewCodeBlock returns a new `CodeBlock`.
func NewCodeBlock(statements ...Statement) *CodeBlock {
	return &CodeBlock{
		Statements: statements,
	}
}

// AddStatements adds statements to `CodeBlock`.
// This method returns a *new* `CodeBlock`; it means this method acts as immutable.
func (c *CodeBlock) AddStatements(statements ...Statement) *CodeBlock {
	return &CodeBlock{
		Statements: append(c.Statements, statements...),
	}
}

// Generate generates plain code block as golang code.
func (c *CodeBlock) Generate(indentLevel int) (string, error) {
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
