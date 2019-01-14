package generator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldPackageStringifyBeSucceeded(t *testing.T) {
	packageName := "foobar"
	packageComponent := NewPackage(packageName)

	gen, err := packageComponent.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("package %s\n", packageName), gen)
}
