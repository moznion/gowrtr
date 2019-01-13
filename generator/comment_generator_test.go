package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCommentStatement(t *testing.T) {
	{
		generator := NewCommentGenerator("this is a comment")
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, "//this is a comment\n", gen)
	}

	{
		generator := NewCommentGenerator(" this is a comment")
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, "\t\t// this is a comment\n", gen)
	}
}
