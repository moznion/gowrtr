package generator

// RawStatement represents a code generator for `raw statement`.
// `raw statement` means plain text statement.
type RawStatement struct {
	Statement   string
	WithNewline bool
}

// NewRawStatement returns a new `RawStatement`.
// `withNewline` is an option to control whether append a newline to statement or not. The default value is `true`.
func NewRawStatement(stmt string, withNewline ...bool) *RawStatement {
	nl := true
	if len(withNewline) > 0 {
		nl = withNewline[0]
	}

	return &RawStatement{
		Statement:   stmt,
		WithNewline: nl,
	}
}

// Generate generates a raw statement.
func (r *RawStatement) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	newline := ""
	if r.WithNewline {
		newline = "\n"
	}

	return indent + r.Statement + newline, nil
}
