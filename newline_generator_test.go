package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateNewlineSuccessful(t *testing.T) {
	generator := NewNewlineGenerator()
	gen, err := generator.GenerateCode(0)
	assert.NoError(t, err)
	assert.Equal(t, "\n", gen)
}
