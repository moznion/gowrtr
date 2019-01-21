package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGeneratingFuncSignatureBeSuccessful(t *testing.T) {
	dataset := map[string]*FuncSignature{
		"myFunc()": NewFuncSignature(
			"myFunc",
		),

		"myFunc(foo string)": NewFuncSignature(
			"myFunc",
		).AddParameters(NewFuncParameter("foo", "string")),

		"myFunc(foo string, bar int)": NewFuncSignature(
			"myFunc",
		).AddParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int"),
		),

		"myFunc(foo, bar string)": NewFuncSignature(
			"myFunc",
		).AddParameters(
			NewFuncParameter("foo", ""),
			NewFuncParameter("bar", "string"),
		),

		"myFunc(foo string, bar int) string": NewFuncSignature(
			"myFunc",
		).AddParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int"),
		).AddReturnTypes("string"),

		"myFunc(foo string, bar int) (string, error)": NewFuncSignature(
			"myFunc",
		).AddParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int"),
		).AddReturnTypes("string", "error"),

		"myFunc(buz error) int64": NewFuncSignature(
			"myFunc",
		).AddParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int"),
		).AddReturnTypes("string", "error").
			Parameters(NewFuncParameter("buz", "error")).
			ReturnTypes("int64"),
	}

	for expected, signature := range dataset {
		gen, err := signature.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldRaiseErrorWhenFuncNameIsEmpty(t *testing.T) {
	sig := NewFuncSignature("")

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}

func TestShouldRaiseErrorWhenFuncParameterNameIsEmpty(t *testing.T) {
	sig := NewFuncSignature("myFunc").AddParameters(
		NewFuncParameter("foo", "string"),
		NewFuncParameter("", "int"),
		NewFuncParameter("buz", "error"),
	)

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.FuncParameterNameIsEmptyErr().Error())
}

func TestShouldRaiseErrorWhenLastFuncParameterTypeIsEmpty(t *testing.T) {
	sig := NewFuncSignature("myFunc").AddParameters(
		NewFuncParameter("foo", "string"),
		NewFuncParameter("bar", ""),
		NewFuncParameter("buz", ""),
	)

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.LastFuncParameterTypeIsEmptyErr().Error())
}

func TestShouldGeneratingFuncSignatureWithNamedReturnValue(t *testing.T) {
	{
		sig, err := NewFuncSignature("myFunc").ReturnTypes("err error").Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, "myFunc() (err error)", sig)
	}

	{
		sig, err := NewFuncSignature("myFunc").ReturnTypes("s string", "err error").Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, "myFunc() (s string, err error)", sig)
	}
}
