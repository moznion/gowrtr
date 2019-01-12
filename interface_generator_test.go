package gowrtr

import (
	"testing"

	"github.com/moznion/gowrtr/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGeneratingInterfaceCodeBeSuccessful(t *testing.T) {
	exp1 := `type myInterface interface {
}
`
	exp2 := `type myInterface interface {
	myFunc()
}
`
	exp3 := `type myInterface interface {
	myFunc1()
	myFunc2(foo string) (string, error)
}
`
	dataset := map[string]*InterfaceGenerator{
		exp1: NewInterfaceGenerator("myInterface"),
		exp2: NewInterfaceGenerator("myInterface").
			AddFuncSignature(NewFuncSignatureGenerator("myFunc")),
		exp3: NewInterfaceGenerator(
			"myInterface",
			NewFuncSignatureGenerator("myFunc1"),
		).AddFuncSignature(
			NewFuncSignatureGenerator("myFunc2").
				AddFuncParameters(NewFuncParameter("foo", "string")).
				AddReturnTypes("string", "error"),
		),
	}

	for expected, in := range dataset {
		got, err := in.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	}
}

func TestShouldGeneratingInterfaceCodeWithIndentBeSuccessful(t *testing.T) {
	exp1 := `		type myInterface interface {
		}
`
	exp2 := `		type myInterface interface {
			myFunc()
		}
`
	exp3 := `		type myInterface interface {
			myFunc1()
			myFunc2(foo string) (string, error)
		}
`
	dataset := map[string]*InterfaceGenerator{
		exp1: NewInterfaceGenerator("myInterface"),
		exp2: NewInterfaceGenerator("myInterface").
			AddFuncSignature(NewFuncSignatureGenerator("myFunc")),
		exp3: NewInterfaceGenerator(
			"myInterface",
			NewFuncSignatureGenerator("myFunc1"),
		).AddFuncSignature(
			NewFuncSignatureGenerator("myFunc2").
				AddFuncParameters(NewFuncParameter("foo", "string")).
				AddReturnTypes("string", "error"),
		),
	}

	for expected, in := range dataset {
		got, err := in.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	}
}

func TestShouldRaiseErrorWhenInterfaceNameIsEmpty(t *testing.T) {
	in := NewInterfaceGenerator("")
	_, err := in.Generate(0)
	assert.EqualError(t, err, errmsg.InterfaceNameIsEmptyError().Error())
}

func TestShouldRaiseErrorWhenFuncSignatureRaisesError(t *testing.T) {
	in := NewInterfaceGenerator(
		"myInterface",
		NewFuncSignatureGenerator(""),
	)
	_, err := in.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
