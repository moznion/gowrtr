package gowrtr

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
