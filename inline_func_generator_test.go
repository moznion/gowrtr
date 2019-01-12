package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateInlineFunc(t *testing.T) {
	generator := NewInlineFuncGenerator(
		false,
		NewInlineFuncSignatureGenerator(),
	)

	{
		expected := `func() {
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		expected := `func() {
	// do something
	fmt.Printf("%d", i)
}
`
		generator = generator.AddStatements(
			NewCommentGenerator(" do something"),
			NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
		)
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateInlineFuncWithSignature(t *testing.T) {
	generator := NewInlineFuncGenerator(
		false,
		NewInlineFuncSignatureGenerator().
			AddFuncParameters(
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int64"),
			).
			AddReturnTypes("string", "error"),
		NewCommentGenerator(" do something"),
		NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
	)

	expected := `func(foo string, bar int64) (string, error) {
	// do something
	fmt.Printf("%d", i)
}
`
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateInlineGoFuncWithInvocation(t *testing.T) {
	generator := NewInlineFuncGenerator(
		true,
		NewInlineFuncSignatureGenerator().
			AddFuncParameters(
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int64"),
			).
			AddReturnTypes("string", "error"),
		NewCommentGenerator(" do something"),
		NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
	).AddFuncInvocation(NewFuncInvocationGenerator("foo", "bar"))

	expected := `go func(foo string, bar int64) (string, error) {
	// do something
	fmt.Printf("%d", i)
}(foo, bar)
`
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}
