package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateIfCode(t *testing.T) {
	generator := NewIfGenerator("i > 0",
		NewCommentGenerator(" do something"),
		NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
	)

	{
		expected := `if i > 0 {
	// do something
	fmt.Printf("%d", i)
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		expected := `		if i > 0 {
			// do something
			fmt.Printf("%d", i)
		}
`
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateIfCodeWithExpandingMethod(t *testing.T) {
	generator := NewIfGenerator("i > 0").
		AddStatements(
			NewCommentGenerator(" XXX: test test"),
			NewCommentGenerator(" do something"),
		).
		AddStatements(NewRawStatementGenerator(`fmt.Printf("%d", i)`, true))

	expected := `if i > 0 {
	// XXX: test test
	// do something
	fmt.Printf("%d", i)
}
`
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateIfCodeGiveUpWhenStatementGeneratorRaisesError(t *testing.T) {
	generator := NewIfGenerator(
		"i > 0",
		NewFuncGenerator(nil, NewFuncSignatureGenerator("")),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
