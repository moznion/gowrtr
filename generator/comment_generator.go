package gowrtr

import "fmt"

type CommentGenerator struct {
	Comment string
}

func NewCommentGenerator(comment string) *CommentGenerator {
	return &CommentGenerator{
		Comment: comment,
	}
}

func (c *CommentGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)
	return fmt.Sprintf("%s//%s\n", indent, c.Comment), nil
}
