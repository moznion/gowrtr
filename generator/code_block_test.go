package generator

import (
	"regexp"
	"strings"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCodeBlock(t *testing.T) {
	generator := NewCodeBlock(
		NewComment(" do something"),
		NewRawStatement(`fmt.Printf("%d", i)`),
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

	{
		generator = generator.Statements(NewComment("modified"))
		expected := `{
	//modified
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateCodeBlockWithEmpty(t *testing.T) {
	generator := NewCodeBlock()

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
	generator := NewCodeBlock().AddStatements(
		NewComment(" XXX: test test"),
		NewComment(" do something"),
	).AddStatements(
		NewRawStatement(`fmt.Printf("%d", i)`),
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

func TestShouldGenerateCodeBlockGiveUpWhenStatementRaisesError(t *testing.T) {
	generator := NewCodeBlock(
		NewFunc(nil, NewFuncSignature("")),
	)
	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncNameIsEmptyError("").Error(), " ")[0],
	), err.Error())
}
