package generator

import "fmt"

// Comment represents a code generator for one line comment.
type Comment struct {
	Comment string
}

// NewComment returns a new `Comment`.
func NewComment(comment string) *Comment {
	return &Comment{
		Comment: comment,
	}
}

// Generate generates one line comment statement.
func (c *Comment) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)
	return fmt.Sprintf("%s//%s\n", indent, c.Comment), nil
}
