package generator

// Statement is an interface that has a responsibility to generate the golang code.
type Statement interface {
	Generate(indentLevel int) (string, error)
}

func BuildIndent(indentLevel int) string {
	indent := ""
	for i := 0; i < indentLevel; i++ {
		indent += "\t"
	}
	return indent
}
