package generator

// CodeBlock represents a code generator for plain code block.
//
// example:
// {
//   // do something
//   fmt.Println("blah blah")
// }
type CodeBlock struct {
	statements []Statement
}

// NewCodeBlock returns a new `CodeBlock`.
func NewCodeBlock(statements ...Statement) *CodeBlock {
	return &CodeBlock{
		statements: statements,
	}
}

// AddStatements adds statements to `CodeBlock`. This does *not* set, just add.
// This method returns a *new* `CodeBlock`; it means this method acts as immutable.
func (c *CodeBlock) AddStatements(statements ...Statement) *CodeBlock {
	return &CodeBlock{
		statements: append(c.statements, statements...),
	}
}

// Statements sets statements to `CodeBlock`. This does *not* add, just set.
// This method returns a *new* `CodeBlock`; it means this method acts as immutable.
func (c *CodeBlock) Statements(statements ...Statement) *CodeBlock {
	return &CodeBlock{
		statements: statements,
	}
}

// Generate generates plain code block as golang code.
func (c *CodeBlock) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := indent + "{\n"

	nextIndentLevel := indentLevel + 1
	for _, generator := range c.statements {
		gen, err := generator.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += indent + "}\n"
	return stmt, nil
}
