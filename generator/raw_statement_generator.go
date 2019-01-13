package generator

// RawStatementGenerator represents a code generator for `raw statement`.
// `raw statement` means plain text statement.
type RawStatementGenerator struct {
	Statement   string
	WithNewline bool
}

// NewRawStatementGenerator returns a new `RawStatementGenerator`.
// `withNewline` is an option to control whether append a newline to statement or not. The default value is `true`.
func NewRawStatementGenerator(stmt string, withNewline ...bool) *RawStatementGenerator {
	nl := true
	if len(withNewline) > 0 {
		nl = withNewline[0]
	}

	return &RawStatementGenerator{
		Statement:   stmt,
		WithNewline: nl,
	}
}

// Generate generates a raw statement.
func (r *RawStatementGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	newline := ""
	if r.WithNewline {
		newline = "\n"
	}

	return indent + r.Statement + newline, nil
}
