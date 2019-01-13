package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateImportStatementBeSucceeded(t *testing.T) {
	importGenerator := NewImportGenerator("fmt", "", "math").AddImports("os")

	expected := `import (
	"fmt"
	"math"
	"os"
)
`

	gen, err := importGenerator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateImportStatementBeSucceededWithSingleImportee(t *testing.T) {
	importGenerator := NewImportGenerator().AddImports("fmt")

	expected := `import (
	"fmt"
)
`

	gen, err := importGenerator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateImportStatementBeEmpty(t *testing.T) {
	importGenerator := NewImportGenerator()

	gen, err := importGenerator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "", gen)
}
