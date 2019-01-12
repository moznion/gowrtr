package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateRawStatementSuccessful(t *testing.T) {
	generator := NewRawStatementGenerator(`i := 0`)

	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "i := 0", gen)

	gen, err = generator.Generate(2)
	assert.NoError(t, err)
	assert.Equal(t, "\t\ti := 0", gen)
}

func TestShouldGenerateRawStatementWithNewlineOption(t *testing.T) {
	{
		generator := NewRawStatementGenerator(`i := 0`, true)

		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, "i := 0\n", gen)

		gen, err = generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, "\t\ti := 0\n", gen)
	}

	{
		generator := NewRawStatementGenerator(`i := 0`, false)

		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, "i := 0", gen)

		gen, err = generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, "\t\ti := 0", gen)
	}
}
