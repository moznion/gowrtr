package gowrtr

type RawStatementGenerator struct {
	Statement string
}

func NewRawStatementGenerator(stmt string) *RawStatementGenerator {
	return &RawStatementGenerator{
		Statement: stmt,
	}
}

func (r *RawStatementGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)
	return indent + r.Statement + "\n", nil
}
