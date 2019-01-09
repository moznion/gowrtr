package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateImportStatementBeSucceeded(t *testing.T) {
	importComponent := NewImport("fmt", "math", "os")

	expected := `import (
	"fmt"
	"math"
	"os"
)`

	gen, err := importComponent.GenerateCode()
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateImportStatementBeSucceededWithSingleImportee(t *testing.T) {
	importComponent := NewImport("fmt")

	expected := `import (
	"fmt"
)`

	gen, err := importComponent.GenerateCode()
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateImportStatementBeEmpty(t *testing.T) {
	importComponent := NewImport()

	gen, err := importComponent.GenerateCode()
	assert.NoError(t, err)
	assert.Equal(t, "", gen)
}
