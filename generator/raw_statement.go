package generator

// RawStatement represents a code generator for `raw statement`.
// `raw statement` means plain text statement.
type RawStatement struct {
	Statement   string
	WithNewline bool
}

// NewRawStatement returns a new `RawStatement`.
func NewRawStatement(stmt string) *RawStatement {
	return &RawStatement{
		Statement:   stmt,
		WithNewline: true,
	}
}

// WithNewLine specifies whether append newline or not.
// Default value is `true`, so this method might be used when you want to suppress to break the line.
func (r *RawStatement) WithNewLine(with bool) *RawStatement {
	return &RawStatement{
		Statement:   r.Statement,
		WithNewline: with,
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
