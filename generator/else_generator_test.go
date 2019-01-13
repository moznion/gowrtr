package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateElseCode(t *testing.T) {
	generator := NewElseGenerator(
		NewCommentGenerator(" do something"),
		NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
	)

	{
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := ` else {
	// do something
	fmt.Printf("%d", i)
}`
		assert.Equal(t, expected, gen)
	}

	{
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		expected := ` else {
			// do something
			fmt.Printf("%d", i)
		}`
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateElseCodeRaisesError(t *testing.T) {
	generator := NewElseGenerator(
		NewFuncGenerator(nil, NewFuncSignatureGenerator("")),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
