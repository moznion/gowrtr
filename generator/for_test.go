package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"
	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateForCode(t *testing.T) {
	generator := NewFor(
		"i := 0; i < foo; i++",
		NewComment(" do something"),
		NewRawStatement(`fmt.Printf("%d", i)`),
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

	{
		generator = generator.Statements(NewComment("modified"))
		expected := `for i := 0; i < foo; i++ {
	//modified
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateForCodeWithExpandingMethod(t *testing.T) {
	generator := NewFor("i := 0; i < foo; i++").
		AddStatements(
			NewComment(" XXX: test test"),
			NewComment(" do something"),
		).
		AddStatements(NewRawStatement(`fmt.Printf("%d", i)`))

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

func TestShouldGenerateForCodeGiveUpWhenStatementRaisesError(t *testing.T) {
	generator := NewFor(
		"i := 0; i < foo; i++",
		NewFunc(nil, NewFuncSignature("")),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
