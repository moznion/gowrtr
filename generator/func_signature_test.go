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
		).AddFuncParameters(NewFuncParameter("foo", "string")),

		"myFunc(foo string, bar int)": NewFuncSignature(
			"myFunc",
		).AddFuncParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int"),
		),

		"myFunc(foo, bar string)": NewFuncSignature(
			"myFunc",
		).AddFuncParameters(
			NewFuncParameter("foo", ""),
			NewFuncParameter("bar", "string"),
		),

		"myFunc(foo string, bar int) string": NewFuncSignature(
			"myFunc",
		).AddFuncParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int"),
		).AddReturnTypes("string"),

		"myFunc(foo string, bar int) (string, error)": NewFuncSignature(
			"myFunc",
		).AddFuncParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int"),
		).AddReturnTypes("string", "error"),
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
	sig := NewFuncSignature("myFunc").AddFuncParameters(
		NewFuncParameter("foo", "string"),
		NewFuncParameter("", "int"),
		NewFuncParameter("buz", "error"),
	)

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.FuncParameterNameIsEmptyErr().Error())
}

func TestShouldRaiseErrorWhenLastFuncParameterTypeIsEmpty(t *testing.T) {
	sig := NewFuncSignature("myFunc").AddFuncParameters(
		NewFuncParameter("foo", "string"),
		NewFuncParameter("bar", ""),
		NewFuncParameter("buz", ""),
	)

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.LastFuncParameterTypeIsEmptyErr().Error())
}
