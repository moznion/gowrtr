package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchClientCallerLineShouldBeNG(t *testing.T) {
	caller := fetchClientCallerLine(10000)
	assert.Empty(t, caller)
}

func TestFetchCallerLineAsSliceShouldBeNG(t *testing.T) {
	callers := fetchClientCallerLineAsSlice(1, 10000)
	assert.Len(t, callers, 1)
	assert.Empty(t, callers[0])
}
