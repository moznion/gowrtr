package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateNewlineSuccessful(t *testing.T) {
	generator := NewNewline()
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "\n", gen)
}
