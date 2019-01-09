package gowrtr

import (
	"testing"

	"github.com/moznion/gowrtr/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGeneratingFuncSignatureBeSuccessful(t *testing.T) {
	dataset := map[string]*FuncSignature{
		"myFunc()": NewFuncSignature(
			"myFunc",
			[]*FuncParameter{},
			[]string{},
		),
		"myFunc(foo string)": NewFuncSignature(
			"myFunc",
			[]*FuncParameter{
				NewFuncParameter("foo", "string"),
			},
			[]string{},
		),
		"myFunc(foo string, bar int)": NewFuncSignature(
			"myFunc",
			[]*FuncParameter{
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int"),
			},
			[]string{},
		),
		"myFunc(foo, bar string)": NewFuncSignature(
			"myFunc",
			[]*FuncParameter{
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			},
			[]string{},
		),
		"myFunc(foo string, bar int) string": NewFuncSignature(
			"myFunc",
			[]*FuncParameter{
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int"),
			},
			[]string{"string"},
		),
		"myFunc(foo string, bar int) (string, error)": NewFuncSignature(
			"myFunc",
			[]*FuncParameter{
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int"),
			},
			[]string{"string", "error"},
		),
	}

	for expected, signature := range dataset {
		gen, err := signature.GenerateCode()
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldRaiseErrorWhenFuncNameIsEmpty(t *testing.T) {
	sig := NewFuncSignature("", []*FuncParameter{}, []string{})

	_, err := sig.GenerateCode()
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}

func TestShouldRaiseErrorWhenFuncParameterNameIsEmpty(t *testing.T) {
	sig := NewFuncSignature("myFunc", []*FuncParameter{
		NewFuncParameter("foo", "string"),
		NewFuncParameter("", "int"),
		NewFuncParameter("buz", "error"),
	}, []string{})

	_, err := sig.GenerateCode()
	assert.EqualError(t, err, errmsg.FuncParameterNameIsEmptyErr().Error())
}

func TestShouldRaiseErrorWhenLastFuncParameterTypeIsEmpty(t *testing.T) {
	sig := NewFuncSignature("myFunc", []*FuncParameter{
		NewFuncParameter("foo", "string"),
		NewFuncParameter("bar", ""),
		NewFuncParameter("buz", ""),
	}, []string{})

	_, err := sig.GenerateCode()
	assert.EqualError(t, err, errmsg.LastFuncParameterTypeIsEmptyErr().Error())
}
