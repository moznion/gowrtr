package generator

import (
	"regexp"
	"strings"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateFuncInvocationCode(t *testing.T) {
	generator := NewFuncInvocation()

	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "()", gen)

	generator = generator.AddParameters("foo")
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(foo)", gen)

	generator = generator.AddParameters("bar")
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(foo, bar)", gen)

	generator = generator.Parameters("buz")
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(buz)", gen)

	generator = generator.GenericsTypes(TypeArguments{"string"})
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "[string](buz)", gen)

	generator = generator.GenericsTypes(TypeArguments{"string", "int"})
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "[string, int](buz)", gen)
}

func TestShouldGenerateFuncInvocationRaisesErrorWhenParameterIsEmpty(t *testing.T) {
	generator := NewFuncInvocation("foo", "", "bar")
	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncInvocationParameterIsEmptyError("").Error(), " ")[0],
	), err.Error())
}
