package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateElseIfCode(t *testing.T) {
	generator := NewElseIf("i > 0",
		NewComment(" do something"),
		NewRawStatement(`fmt.Printf("%d", i)`),
	)

	{
		expected := ` else if i > 0 {
	// do something
	fmt.Printf("%d", i)
}`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		expected := ` else if i > 0 {
			// do something
			fmt.Printf("%d", i)
		}`
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateElseIfWithExpandingMethod(t *testing.T) {
	generator := NewElseIf("i > 0").
		AddStatements(
			NewComment(" XXX: test test"),
			NewComment(" do something"),
		).
		AddStatements(NewRawStatement(`fmt.Printf("%d", i)`))

	expected := ` else if i > 0 {
	// XXX: test test
	// do something
	fmt.Printf("%d", i)
}`
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateElseIfRaisesError(t *testing.T) {
	generator := NewElseIf(
		"i > 0",
		NewFunc(nil, NewFuncSignature("")),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
