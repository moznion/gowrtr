package generator

import "fmt"

// CommentGenerator represents a code generator for one line comment.
type CommentGenerator struct {
	Comment string
}

// NewCommentGenerator returns a new `CommentGenerator`.
func NewCommentGenerator(comment string) *CommentGenerator {
	return &CommentGenerator{
		Comment: comment,
	}
}

// Generate generates one line comment statement.
func (c *CommentGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)
	return fmt.Sprintf("%s//%s\n", indent, c.Comment), nil
}
