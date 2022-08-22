package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateTypeParametersSuccessfully(t *testing.T) {
	stmt, err := TypeParameters{NewTypeParameter("T", "string")}.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "[T string]", stmt)

	stmt, err = TypeParameters{NewTypeParameter("T", "string"), NewTypeParameter("U", "int")}.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "[T string, U int]", stmt)
}

func TestShouldFailToGenerateTypeParametersWhenParameterIsEmpty(t *testing.T) {
	_, err := TypeParameters{NewTypeParameter("", "string")}.Generate(0)
	assert.Error(t, err)
	assert.Equal(t, errmsg.TypeParameterParameterIsEmptyErrType, errmsg.IdentifyErrs(err))
}

func TestShouldGenerateUnionTypeParameterSuccessfully(t *testing.T) {
	stmt, err := TypeParameters{NewTypeParameters("T", []string{"int", "uint"})}.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "[T int | uint]", stmt)
}

func TestShouldFailToGenerateTypeParametersWhenTypeIsEmpty(t *testing.T) {
	_, err := TypeParameters{NewTypeParameter("T", "")}.Generate(0)
	assert.Error(t, err)
	assert.Equal(t, errmsg.TypeParameterTypeIsEmptyErrType, errmsg.IdentifyErrs(err))
}

func TestShouldGenerateUnionTypeParametersWhenTypesSliceIsEmpty(t *testing.T) {
	_, err := TypeParameters{NewTypeParameters("T", []string{})}.Generate(0)
	assert.Error(t, err)
	assert.Equal(t, errmsg.TypeParameterTypeIsEmptyErrType, errmsg.IdentifyErrs(err))
}

func TestShouldGenerateTypeArgumentsSuccessfully(t *testing.T) {
	stmt, err := TypeArguments{"string"}.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "[string]", stmt)

	stmt, err = TypeArguments{"string", "int"}.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, "[string, int]", stmt)
}
