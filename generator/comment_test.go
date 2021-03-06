package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCommentStatement(t *testing.T) {
	{
		generator := NewComment("this is a comment")
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, "//this is a comment\n", gen)
	}

	{
		generator := NewComment(" this is a comment")
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, "\t\t// this is a comment\n", gen)
	}
}

func TestShouldGenerateCommentStatementWithFormatting(t *testing.T) {
	generator := NewCommentf("this is a %s", "comment")
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "//this is a comment\n", gen)
}
