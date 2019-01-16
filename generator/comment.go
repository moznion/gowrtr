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

// NewCommentf returns a new `Comment` with formatting.
// If `args` is not empty, this method formats `stmt` with `args` by `fmt.Sprintf`.
func NewCommentf(comment string, args ...interface{}) *Comment {
	return &Comment{
		Comment: fmt.Sprintf(comment, args...),
	}
}

// Generate generates one line comment statement.
func (c *Comment) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)
	return fmt.Sprintf("%s//%s\n", indent, c.Comment), nil
}
