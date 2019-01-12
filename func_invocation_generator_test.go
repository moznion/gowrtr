package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateFuncInvocationCode(t *testing.T) {
	generator := NewFuncInvocationGenerator()

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
}
