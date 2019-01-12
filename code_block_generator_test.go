package gowrtr

import (
	"testing"

	"github.com/moznion/gowrtr/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCodeBlock(t *testing.T) {
	generator := NewCodeBlockGenerator(
		NewCommentGenerator(" do something"),
		NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
	)

	{
		expected := `{
	// do something
	fmt.Printf("%d", i)
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		expected := `		{
			// do something
			fmt.Printf("%d", i)
		}
`
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateCodeBlockWithEmpty(t *testing.T) {
	generator := NewCodeBlockGenerator()

	{
		expected := `{
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		expected := `		{
		}
`
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateCodeBlockWithExpandingMethod(t *testing.T) {
	generator := NewCodeBlockGenerator().AddStatements(
		NewCommentGenerator(" XXX: test test"),
		NewCommentGenerator(" do something"),
	).AddStatements(
		NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
	)

	expected := `{
	// XXX: test test
	// do something
	fmt.Printf("%d", i)
}
`
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateCodeBlockGiveUpWhenStatementGeneratorRaisesError(t *testing.T) {
	generator := NewCodeBlockGenerator(
		NewFuncGenerator(nil, NewFuncSignatureGenerator("")),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
