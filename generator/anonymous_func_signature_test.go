package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateAnonymousFuncSignatureCode(t *testing.T) {
	generator := NewAnonymousFuncSignature()
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "()", gen)

	generator = generator.
		AddFuncParameters(NewFuncParameter("foo", "string")).
		AddReturnTypes("string")
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(foo string) string", gen)

	generator = NewAnonymousFuncSignature().
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int64"),
		).
		AddReturnTypes("string", "error")
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(foo string, bar int64) (string, error)", gen)
}

func TestShouldGenerateAnonymousFuncSignatureRaisesErrorWhenParamNameIsEmpty(t *testing.T) {
	generator := NewAnonymousFuncSignature().AddFuncParameters(
		NewFuncParameter("", "string"),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncParameterNameIsEmptyErr().Error())
}

func TestShouldGenerateAnonymousFuncSignatureRaisesErrorWhenParamTypeIsEmpty(t *testing.T) {
	generator := NewAnonymousFuncSignature().AddFuncParameters(
		NewFuncParameter("foo", ""),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.LastFuncParameterTypeIsEmptyErr().Error())
}
