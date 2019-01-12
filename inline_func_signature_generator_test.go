package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateInlineFuncSignatureCode(t *testing.T) {
	generator := NewInlineFuncSignatureGenerator()
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "()", gen)

	generator = generator.
		AddFuncParameters(NewFuncParameter("foo", "string")).
		AddReturnTypes("string")
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(foo string) string", gen)

	generator = NewInlineFuncSignatureGenerator().
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int64"),
		).
		AddReturnTypes("string", "error")
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(foo string, bar int64) (string, error)", gen)
}
