package gowrtr

import (
	"testing"

	"github.com/moznion/gowrtr/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGeneratingFuncSignatureBeSuccessful(t *testing.T) {
	dataset := map[string]*FuncSignatureGenerator{
		"myFunc()": NewFuncSignatureGenerator(
			"myFunc",
			[]*FuncParameter{},
			[]string{},
		),
		"myFunc(foo string)": NewFuncSignatureGenerator(
			"myFunc",
			[]*FuncParameter{
				NewFuncParameter("foo", "string"),
			},
			[]string{},
		),
		"myFunc(foo string, bar int)": NewFuncSignatureGenerator(
			"myFunc",
			[]*FuncParameter{
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int"),
			},
			[]string{},
		),
		"myFunc(foo, bar string)": NewFuncSignatureGenerator(
			"myFunc",
			[]*FuncParameter{
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			},
			[]string{},
		),
		"myFunc(foo string, bar int) string": NewFuncSignatureGenerator(
			"myFunc",
			[]*FuncParameter{
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int"),
			},
			[]string{"string"},
		),
		"myFunc(foo string, bar int) (string, error)": NewFuncSignatureGenerator(
			"myFunc",
			[]*FuncParameter{
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int"),
			},
			[]string{"string", "error"},
		),
	}

	for expected, signature := range dataset {
		gen, err := signature.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGeneratingFuncSignatureWithExpandMethod(t *testing.T) {
	generator := NewFuncSignatureGenerator(
		"myFunc",
		[]*FuncParameter{},
		[]string{},
	).AddFuncParameters(
		NewFuncParameter("foo", "string"),
		NewFuncParameter("bar", "int64"),
	).AddReturnTypes("string", "error")

	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "myFunc(foo string, bar int64) (string, error)", gen)
}

func TestShouldRaiseErrorWhenFuncNameIsEmpty(t *testing.T) {
	sig := NewFuncSignatureGenerator("", []*FuncParameter{}, []string{})

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}

func TestShouldRaiseErrorWhenFuncParameterNameIsEmpty(t *testing.T) {
	sig := NewFuncSignatureGenerator("myFunc", []*FuncParameter{
		NewFuncParameter("foo", "string"),
		NewFuncParameter("", "int"),
		NewFuncParameter("buz", "error"),
	}, []string{})

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.FuncParameterNameIsEmptyErr().Error())
}

func TestShouldRaiseErrorWhenLastFuncParameterTypeIsEmpty(t *testing.T) {
	sig := NewFuncSignatureGenerator("myFunc", []*FuncParameter{
		NewFuncParameter("foo", "string"),
		NewFuncParameter("bar", ""),
		NewFuncParameter("buz", ""),
	}, []string{})

	_, err := sig.Generate(0)
	assert.EqualError(t, err, errmsg.LastFuncParameterTypeIsEmptyErr().Error())
}
