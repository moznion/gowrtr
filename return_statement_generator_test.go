package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateReturnStatement(t *testing.T) {
	generator := NewReturnStatementGenerator()
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "return", gen)

	generator = generator.AddReturnItems("foo", "err")
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "return foo, err", gen)
}
