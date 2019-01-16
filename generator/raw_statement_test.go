package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateRawStatementSuccessful(t *testing.T) {
	generator := NewRawStatement(`i := 0`)

	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "i := 0\n", gen)

	gen, err = generator.Generate(2)
	assert.NoError(t, err)
	assert.Equal(t, "\t\ti := 0\n", gen)
}

func TestShouldGenerateRawStatementWithFormattingSuccessful(t *testing.T) {
	generator := NewRawStatementf(`s := "%s"`, "test-str")

	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "s := \"test-str\"\n", gen)
}

func TestShouldGenerateRawStatementWithNewlineOption(t *testing.T) {
	{
		generator := NewRawStatement(`i := 0`).WithNewline(true)

		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, "i := 0\n", gen)

		gen, err = generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, "\t\ti := 0\n", gen)
	}

	{
		generator := NewRawStatement(`i := 0`).WithNewline(false)

		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, "i := 0", gen)

		gen, err = generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, "\t\ti := 0", gen)
	}
}
