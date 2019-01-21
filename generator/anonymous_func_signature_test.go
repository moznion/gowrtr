package generator

import (
	"regexp"
	"strings"
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
		AddParameters(NewFuncParameter("foo", "string")).
		AddReturnTypes("string")
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(foo string) string", gen)

	generator = NewAnonymousFuncSignature().
		AddParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int64"),
		).
		AddReturnTypes("string", "error")
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(foo string, bar int64) (string, error)", gen)

	gen, err = generator.Parameters(NewFuncParameter("buz", "error")).ReturnTypes("error").Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "(buz error) error", gen)
}

func TestShouldGenerateAnonymousFuncSignatureRaisesErrorWhenParamNameIsEmpty(t *testing.T) {
	generator := NewAnonymousFuncSignature().AddParameters(
		NewFuncParameter("", "string"),
	)
	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncParameterNameIsEmptyErr("").Error(), " ")[0],
	), err.Error())
}

func TestShouldGenerateAnonymousFuncSignatureRaisesErrorWhenParamTypeIsEmpty(t *testing.T) {
	generator := NewAnonymousFuncSignature().AddParameters(
		NewFuncParameter("foo", ""),
	)
	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.LastFuncParameterTypeIsEmptyErr("").Error(), " ")[0],
	), err.Error())
}
