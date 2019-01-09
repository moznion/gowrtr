package gowrtr

import (
	"testing"

	"github.com/moznion/gowrtr/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGeneratingInterfaceCodeBeSuccessful(t *testing.T) {
	exp1 := `type myInterface interface {
}`
	exp2 := `type myInterface interface {
	myFunc()
}`
	exp3 := `type myInterface interface {
	myFunc1()
	myFunc2(foo string) (string, error)
}`
	dataset := map[string]*Interface{
		exp1: NewInterface(
			"myInterface",
			[]*FuncSignature{},
		),
		exp2: NewInterface(
			"myInterface",
			[]*FuncSignature{
				NewFuncSignature("myFunc", []*FuncParameter{}, []string{}),
			},
		),
		exp3: NewInterface(
			"myInterface",
			[]*FuncSignature{
				NewFuncSignature("myFunc1", []*FuncParameter{}, []string{}),
				NewFuncSignature(
					"myFunc2",
					[]*FuncParameter{NewFuncParameter("foo", "string")},
					[]string{"string", "error"},
				),
			},
		),
	}

	for expected, in := range dataset {
		got, err := in.GenerateCode()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	}
}

func TestShouldRaiseErrorWhenInterfaceNameIsEmpty(t *testing.T) {
	in := NewInterface("", []*FuncSignature{})
	_, err := in.GenerateCode()
	assert.EqualError(t, err, errmsg.InterfaceNameIsEmptyError().Error())
}

func TestShouldRaiseErrorWhenFuncSignatureRaisesError(t *testing.T) {
	in := NewInterface("myInterface", []*FuncSignature{
		NewFuncSignature("", []*FuncParameter{}, []string{}),
	})
	_, err := in.GenerateCode()
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
