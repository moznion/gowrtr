package gowrtr

import (
	"fmt"
	"testing"
)

func TestShouldPackageStringifyBeSucceeded(t *testing.T) {
	packageName := "foobar"
	packageComponent := NewPackage(packageName)
	if gen := packageComponent.String(); gen != fmt.Sprintf("package %s", packageName) {
		t.Errorf("got unexpected generated code: %s", gen)
	}
}
