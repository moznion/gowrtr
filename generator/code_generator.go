package generator

// StatementGenerator is an interface that has a responsibility to generate the golang's code.
type StatementGenerator interface {
	Generate(indentLevel int) (string, error)
}

func buildIndent(indentLevel int) string {
	indent := ""
	for i := 0; i < indentLevel; i++ {
		indent += "\t"
	}
	return indent
}
