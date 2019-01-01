package gowrtr

import (
	"testing"
)

func TestShouldGenerateImportStatementBeSucceeded(t *testing.T) {
	importComponent := NewImport("fmt", "math", "os")

	expected := `import (
	"fmt"
	"math"
	"os"
)`
	if gen := importComponent.String(); gen != expected {
		t.Errorf("got unexpected generated code:\n%s", gen)
	}
}

func TestShouldGenerateImportStatementBeSucceededWithSingleImportee(t *testing.T) {
	importComponent := NewImport("fmt")

	expected := `import (
	"fmt"
)`
	if gen := importComponent.String(); gen != expected {
		t.Errorf("got unexpected generated code:\n%s", gen)
	}
}

func TestShouldGenerateImportStatementBeEmpty(t *testing.T) {
	importComponent := NewImport()

	if importComponent.String() != "" {
		t.Errorf("generated code is not empty")
	}
}
