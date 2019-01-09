package gowrtr

import (
	"testing"

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

	structComponent, err := NewStruct(structName, fields)
	if err != nil {
		t.Fatalf("unexpected error has come: %s", err)
	}

	expected := "type TestStruct struct {\n" +
		"	Foo string\n" +
		"	Bar int64 `json:\"bar\"`\n" +
		"	buz []byte\n" +
		"}"
	if gen := structComponent.String(); gen != expected {
		t.Errorf("got unexpected generated code: %s", gen)
	}
}

func TestShouldRaiseErrorWhenStructNameIsEmpty(t *testing.T) {
	_, err := NewStruct("", []*StructField{})

	expectedErr := errmsg.StructNameIsNilErr()
	if err == nil || err.Error() != expectedErr.Error() {
		t.Fatalf(`got unexpected error: got="%s", expected="%s"`, err, expectedErr)
	}
}

func TestShouldRaiseErrorWhenFieldNameIsEmpty(t *testing.T) {
	_, err := NewStruct("TestStruct", []*StructField{
		{
			Name: "",
			Type: "string",
		},
	})

	expectedErr := errmsg.StructFieldNameIsEmptyErr()
	if err == nil || err.Error() != expectedErr.Error() {
		t.Fatalf(`got unexpected error: got="%s", expected="%s"`, err, expectedErr)
	}
}

func TestShouldRaiseErrorWhenFieldTypeIsEmpty(t *testing.T) {
	_, err := NewStruct("TestStruct", []*StructField{
		{
			Name: "Foo",
			Type: "",
		},
	})

	expectedErr := errmsg.StructFieldTypeIsEmptyErr()
	if err == nil || err.Error() != expectedErr.Error() {
		t.Fatalf(`got unexpected error: got="%s", expected="%s"`, err, expectedErr)
	}
}
