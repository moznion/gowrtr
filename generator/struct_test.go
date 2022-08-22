package generator

import (
	"regexp"
	"strings"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateStructStatementBeSucceeded(t *testing.T) {
	structName := "TestStruct"

	structGenerator := NewStruct(structName).
		AddField("Foo", "string").
		AddField("Bar", "int64", `json:"bar"`).
		AddField("buz", "[]byte")

	{
		gen, err := structGenerator.Generate(0)
		expected := "type TestStruct struct {\n" +
			"	Foo string\n" +
			"	Bar int64 `json:\"bar\"`\n" +
			"	buz []byte\n" +
			"}\n"
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		gen, err := structGenerator.Generate(2)
		expected := "\t\ttype TestStruct struct {\n" +
			"\t\t\tFoo string\n" +
			"\t\t\tBar int64 `json:\"bar\"`\n" +
			"\t\t\tbuz []byte\n" +
			"\t\t}\n"
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateStructStatementWithTypeParameterSuccessfully(t *testing.T) {
	generator := NewStruct("TestStruct")
	generator = generator.
		TypeParameters(TypeParameters{NewTypeParameter("T", "string")}).
		AddField("Foo", "T").
		AddField("Bar", "int64", `json:"bar"`)

	gen, err := generator.Generate(0)
	expected := "type TestStruct[T string] struct {\n" +
		"	Foo T\n" +
		"	Bar int64 `json:\"bar\"`\n" +
		"}\n"
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateStructStatementWithTypeParametersSuccessfully(t *testing.T) {
	generator := NewStruct("TestStruct")
	generator = generator.
		TypeParameters(TypeParameters{
			NewTypeParameter("T", "string"),
			NewTypeParameter("U", "int64"),
		}).
		AddField("Foo", "T").
		AddField("Bar", "U", `json:"bar"`)

	gen, err := generator.Generate(0)
	expected := "type TestStruct[T string, U int64] struct {\n" +
		"	Foo T\n" +
		"	Bar U `json:\"bar\"`\n" +
		"}\n"
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldRaiseErrorWhenStructNameIsEmpty(t *testing.T) {
	structGenerator := NewStruct("")

	_, err := structGenerator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.StructNameIsNilErr("").Error(), " ")[0],
	), err.Error())
}

func TestShouldRaiseErrorWhenFieldNameIsEmpty(t *testing.T) {
	structGenerator := NewStruct("TestStruct").AddField("", "string")
	_, err := structGenerator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.StructFieldNameIsEmptyErr("").Error(), " ")[0],
	), err.Error())
}

func TestShouldRaiseErrorWhenFieldTypeIsEmpty(t *testing.T) {
	structGenerator := NewStruct("TestStruct").AddField("Foo", "")
	_, err := structGenerator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.StructFieldTypeIsEmptyErr("").Error(), " ")[0],
	), err.Error())
}

func TestShouldRaiseErrorWhenInvalidTypeParameterHasGiven(t *testing.T) {
	generator := NewStruct("TestStruct")
	generator = generator.
		TypeParameters(TypeParameters{NewTypeParameter("", "string")}).
		AddField("Foo", "T")
	_, err := generator.Generate(0)
	assert.Error(t, err)
	assert.Equal(t, errmsg.TypeParameterParameterIsEmptyErrType, errmsg.IdentifyErrs(err))
}
