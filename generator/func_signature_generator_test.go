package gowrtr

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGeneratingFuncSignatureBeSuccessful(t *testing.T) {
	dataset := map[string]*FuncSignatureGenerator{
		"myFunc()": NewFuncSignatureGenerator(
			"myFunc",
		),

		"myFunc(foo string)": NewFuncSignatureGenerator(
			"myFunc",
		).AddFuncParameters(NewFuncParameter("foo", "string")),

		"myFunc(foo string, bar int)": NewFuncSignatureGenerator(
			"myFunc",
		).AddFuncParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int"),
		),

		"myFunc(foo, bar string)": NewFuncSignatureGenerator(
			"myFunc",
		).AddFuncParameters(
			NewFuncParameter("foo", ""),
			NewFuncParameter("bar", "string"),
		),

		"myFunc(foo string, bar int) string": NewFuncSignatureGenerator(
			"myFunc",
		).AddFuncParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int"),
		).AddReturnTypes("string"),

		"myFunc(foo string, bar int) (string, error)": NewFuncSignatureGenerator(
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
	sig := NewFuncSignatureGenerator("")

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}

func TestShouldRaiseErrorWhenFuncParameterNameIsEmpty(t *testing.T) {
	sig := NewFuncSignatureGenerator("myFunc").AddFuncParameters(
		NewFuncParameter("foo", "string"),
		NewFuncParameter("", "int"),
		NewFuncParameter("buz", "error"),
	)

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.FuncParameterNameIsEmptyErr().Error())
}

func TestShouldRaiseErrorWhenLastFuncParameterTypeIsEmpty(t *testing.T) {
	sig := NewFuncSignatureGenerator("myFunc").AddFuncParameters(
		NewFuncParameter("foo", "string"),
		NewFuncParameter("bar", ""),
		NewFuncParameter("buz", ""),
	)

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.LastFuncParameterTypeIsEmptyErr().Error())
}
