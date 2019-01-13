package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateFuncCode(t *testing.T) {
	generator := NewFuncGenerator(
		NewFuncReceiverGenerator("m", "*MyStruct"),
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
		expected := `func (m *MyStruct) myFunc(foo, bar string) (string, error) {
	// do something

	return foo+bar, nil
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		expected := `		func (m *MyStruct) myFunc(foo, bar string) (string, error) {
			// do something

			return foo+bar, nil
		}
`
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateFuncCodeGiveUpWhenFuncNameIsEmpty(t *testing.T) {
	generator := NewFuncGenerator(
		nil,
		NewFuncSignatureGenerator("").
			AddFuncParameters(
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}

func TestShouldGenerateFuncCodeGiveUpWhenFuncSignatureIsNil(t *testing.T) {
	generator := NewFuncGenerator(
		nil,
		nil,
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncSignatureIsNilError().Error())
}

func TestShouldGenerateFuncCodeGiveUpWhenFuncReceiverRaisesError(t *testing.T) {
	generator := NewFuncGenerator(
		NewFuncReceiverGenerator("", "*Foo"),
		NewFuncSignatureGenerator("myFunc").
			AddFuncParameters(
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncReceiverNameIsEmptyError().Error())
}

func TestShouldGenerateFuncCodeGiveUpWhenStatementGeneratorRaisesError(t *testing.T) {
	generator := NewFuncGenerator(
		nil,
		NewFuncSignatureGenerator("myFunc").
			AddFuncParameters(
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			).
			AddReturnTypes("string", "error"),
		NewFuncGenerator(nil, NewFuncSignatureGenerator("")),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
