package generator

import "fmt"

// Comment represents a code generator for one line comment.
type Comment struct {
	comment string
}

// NewComment returns a new `Comment`.
func NewComment(comment string) *Comment {
	return &Comment{
		comment: comment,
	}
}

// NewCommentf returns a new `Comment` with formatting.
// If `args` is not empty, this method formats `stmt` with `args` by `fmt.Sprintf`.
func NewCommentf(comment string, args ...interface{}) *Comment {
	return &Comment{
		comment: fmt.Sprintf(comment, args...),
	}
}

// Generate generates one line comment statement.
func (c *Comment) Generate(indentLevel int) (string, error) {
	indent := BuildIndent(indentLevel)
	return fmt.Sprintf("%s//%s\n", indent, c.comment), nil
}
