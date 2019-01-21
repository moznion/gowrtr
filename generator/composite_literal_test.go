package generator

import (
	"regexp"
	"strings"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCompositeLiteralBeSuccess(t *testing.T) {
	composeGenerator := NewCompositeLiteral("&Struct").
		AddField("foo", NewRawStatement(`"foo-value"`)).
		AddFieldStr("bar", "bar-value").
		AddField("buz", NewAnonymousFunc(
			false,
			NewAnonymousFuncSignature().ReturnTypes("bool"),
			NewReturnStatement("true"),
		).Invocation(NewFuncInvocation())).
		AddFieldRaw("qux", 12345).
		AddFieldRaw("foobar", false)

	{
		gen, err := composeGenerator.Generate(0)

		assert.NoError(t, err)
		expected := `&Struct{
	foo: "foo-value",
	bar: "bar-value",
	buz: func() bool {
		return true
	}(),
	qux: 12345,
	foobar: false,
}
`
		assert.Equal(t, expected, gen)
	}

	{
		gen, err := composeGenerator.Generate(2)

		assert.NoError(t, err)
		expected := `		&Struct{
			foo: "foo-value",
			bar: "bar-value",
			buz: func() bool {
				return true
			}(),
			qux: 12345,
			foobar: false,
		}
`
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateCompositeLiteralWithEmptyKey(t *testing.T) {
	composeGenerator := NewCompositeLiteral("[]string").
		AddFieldStr("", "foo").
		AddFieldStr("", "bar").
		AddFieldStr("", "buz")
	gen, err := composeGenerator.Generate(0)
	expected := `[]string{
	"foo",
	"bar",
	"buz",
}
`
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateCompositeLiteralRaiseError(t *testing.T) {
	_, err := NewCompositeLiteral("").AddField("foo", NewIf("")).Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.IfConditionIsEmptyError("").Error(), " ")[0],
	), err.Error())
}

func TestShouldGenerateCompositeLiteralRaiseErrorWhenValueIsEmpty(t *testing.T) {
	_, err := NewCompositeLiteral("[]string").AddField("foo", NewRawStatement("")).Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.ValueOfCompositeLiteralIsEmptyError("").Error(), " ")[0],
	), err.Error())
}
