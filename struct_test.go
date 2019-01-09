package gowrtr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moznion/gowrtr/errmsg"
)

func TestShouldGenerateStructStatementBeSucceeded(t *testing.T) {
	structName := "TestStruct"
	fields := []*StructField{
		{
			Name: "Foo",
			Type: "string",
		},
		{
			Name: "Bar",
			Type: "int64",
			Tag:  `json:"bar"`,
		},
		{
			Name: "buz",
			Type: "[]byte",
		},
	}

	structComponent := NewStruct(structName, fields)

	expected := "type TestStruct struct {\n" +
		"	Foo string\n" +
		"	Bar int64 `json:\"bar\"`\n" +
		"	buz []byte\n" +
		"}"

	gen, err := structComponent.GenerateCode()
	assert.NoError(t, err)
	assert.Equal(t, gen, expected)
}

func TestShouldRaiseErrorWhenStructNameIsEmpty(t *testing.T) {
	structComponent := NewStruct("", []*StructField{})

	_, err := structComponent.GenerateCode()
	assert.EqualError(t, err, errmsg.StructNameIsNilErr().Error())
}

func TestShouldRaiseErrorWhenFieldNameIsEmpty(t *testing.T) {
	structComponent := NewStruct("TestStruct", []*StructField{
		{
			Name: "",
			Type: "string",
		},
	})

	_, err := structComponent.GenerateCode()
	assert.EqualError(t, err, errmsg.StructFieldNameIsEmptyErr().Error())
}

func TestShouldRaiseErrorWhenFieldTypeIsEmpty(t *testing.T) {
	structComponent := NewStruct("TestStruct", []*StructField{
		{
			Name: "Foo",
			Type: "",
		},
	})

	_, err := structComponent.GenerateCode()
	assert.EqualError(t, err, errmsg.StructFieldTypeIsEmptyErr().Error())
}
