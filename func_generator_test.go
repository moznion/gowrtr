package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateFuncCode(t *testing.T) {
	generator := NewFuncGenerator(
		nil,
		NewFuncSignatureGenerator("myFunc").
			AddFuncParameters(
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			).
			AddReturnTypes("string", "error"),
	).AddStatements(
		NewCommentGenerator(" do something"),
		NewNewlineGenerator(),
		NewReturnStatementGenerator("foo+bar", "nil"),
	)

	{
		expected := `func myFunc(foo, bar string) (string, error) {
	// do something

	return foo+bar, nil
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		expected := `		func myFunc(foo, bar string) (string, error) {
			// do something

			return foo+bar, nil
		}
`
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}
