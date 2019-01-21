package generator

import (
	"regexp"
	"strings"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateFuncCode(t *testing.T) {
	generator := NewFunc(
		NewFuncReceiver("m", "*MyStruct"),
		NewFuncSignature("myFunc").
			AddParameters(
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			).
			AddReturnTypes("string", "error"),
	).AddStatements(
		NewComment(" do something"),
		NewNewline(),
		NewReturnStatement("foo+bar", "nil"),
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

	{
		generator = generator.Statements(NewComment("modified"))
		expected := `func (m *MyStruct) myFunc(foo, bar string) (string, error) {
	//modified
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateFuncCodeGiveUpWhenFuncNameIsEmpty(t *testing.T) {
	generator := NewFunc(
		nil,
		NewFuncSignature("").
			AddParameters(
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			),
	)

	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncNameIsEmptyError("").Error(), " ")[0],
	), err.Error())
}

func TestShouldGenerateFuncCodeGiveUpWhenFuncSignatureIsNil(t *testing.T) {
	generator := NewFunc(
		nil,
		nil,
	)

	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncSignatureIsNilError("").Error(), " ")[0],
	), err.Error())
}

func TestShouldGenerateFuncCodeGiveUpWhenFuncReceiverRaisesError(t *testing.T) {
	generator := NewFunc(
		NewFuncReceiver("", "*Foo"),
		NewFuncSignature("myFunc").
			AddParameters(
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			),
	)

	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncReceiverNameIsEmptyError("").Error(), " ")[0],
	), err.Error())
}

func TestShouldGenerateFuncCodeGiveUpWhenStatementRaisesError(t *testing.T) {
	generator := NewFunc(
		nil,
		NewFuncSignature("myFunc").
			AddParameters(
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			).
			AddReturnTypes("string", "error"),
		NewFunc(nil, NewFuncSignature("")),
	)

	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncNameIsEmptyError("").Error(), " ")[0],
	), err.Error())
}
