package generator

import "fmt"

// RawStatement represents a code generator for `raw statement`.
// `raw statement` means plain text statement.
type RawStatement struct {
	statement   string
	withNewline bool
}

// NewRawStatement returns a new `RawStatement`.
func NewRawStatement(stmt string) *RawStatement {
	return &RawStatement{
		statement:   stmt,
		withNewline: true,
	}
}

// NewRawStatementf returns a new `RawStatement` with formatting.
// If `args` is not empty, this method formats `stmt` with `args` by `fmt.Sprintf`.
func NewRawStatementf(stmt string, args ...interface{}) *RawStatement {
	return &RawStatement{
		statement:   fmt.Sprintf(stmt, args...),
		withNewline: true,
	}
}

// WithNewline specifies whether append newline or not.
// Default value is `true`, so this method might be used when you want to suppress to break the line.
func (r *RawStatement) WithNewline(with bool) *RawStatement {
	return &RawStatement{
		statement:   r.statement,
		withNewline: with,
	}
}

// Generate generates a raw statement.
func (r *RawStatement) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	newline := ""
	if r.withNewline {
		newline = "\n"
	}

	return indent + r.statement + newline, nil
}
