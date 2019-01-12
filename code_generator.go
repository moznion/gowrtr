package gowrtr

type CodeGenerator interface {
	Generate(indentLevel int) (string, error)
}

func buildIndent(indentLevel int) string {
	indent := ""
	for i := 0; i < indentLevel; i++ {
		indent += "\t"
	}
	return indent
}

func Generate(indentLevel int, generators ...CodeGenerator) (string, error) {
	generatedCode := ""

	for _, generator := range generators {
		gen, err := generator.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		generatedCode += gen
	}

	return generatedCode, nil
}
