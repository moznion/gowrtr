package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateForCode(t *testing.T) {
	generator := NewForGenerator(
		"i := 0; i < foo; i++",
		NewCommentGenerator(" do something"),
		NewRawStatementGenerator(`fmt.Printf("%d", i)`),
	)

	{
		expected := `for i := 0; i < foo; i++ {
	// do something
	fmt.Printf("%d", i)
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		expected := `		for i := 0; i < foo; i++ {
			// do something
			fmt.Printf("%d", i)
		}
`
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateForCodeWithExpandingMethod(t *testing.T) {
	generator := NewForGenerator("i := 0; i < foo; i++").
		AddStatements(
			NewCommentGenerator(" XXX: test test"),
			NewCommentGenerator(" do something"),
		).
		AddStatements(NewRawStatementGenerator(`fmt.Printf("%d", i)`))

	expected := `for i := 0; i < foo; i++ {
	// XXX: test test
	// do something
	fmt.Printf("%d", i)
}
`
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}
