package gowrtr

import (
	"testing"

	"github.com/moznion/gowrtr/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateStructStatementBeSucceeded(t *testing.T) {
	structName := "TestStruct"

	structGenerator := NewStructGenerator(structName).
		AddField("Foo", "string").
		AddField("Bar", "int64", `json:"bar"`).
		AddField("buz", "[]byte")

	expected := "type TestStruct struct {\n" +
		"	Foo string\n" +
		"	Bar int64 `json:\"bar\"`\n" +
		"	buz []byte\n" +
		"}"

	gen, err := structGenerator.GenerateCode()
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldRaiseErrorWhenStructNameIsEmpty(t *testing.T) {
	structComponent := NewStructGenerator("")

	_, err := structComponent.GenerateCode()
	assert.EqualError(t, err, errmsg.StructNameIsNilErr().Error())
}

func TestShouldRaiseErrorWhenFieldNameIsEmpty(t *testing.T) {
	structComponent := NewStructGenerator("TestStruct").AddField("", "string")
	_, err := structComponent.GenerateCode()
	assert.EqualError(t, err, errmsg.StructFieldNameIsEmptyErr().Error())
}

func TestShouldRaiseErrorWhenFieldTypeIsEmpty(t *testing.T) {
	structComponent := NewStructGenerator("TestStruct").AddField("Foo", "")
	_, err := structComponent.GenerateCode()
	assert.EqualError(t, err, errmsg.StructFieldTypeIsEmptyErr().Error())
}
