package generator

import (
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
}

func TestShouldGenerateFuncInvocationRaisesErrorWhenParameterIsEmpty(t *testing.T) {
	generator := NewFuncInvocation("foo", "", "bar")
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncInvocationParameterIsEmptyError().Error())
}
