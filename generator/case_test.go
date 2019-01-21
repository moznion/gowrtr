package generator

import (
	"regexp"
	"strings"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCase(t *testing.T) {
	generator := NewCase(
		`"foo"`,
		NewComment(" XXX test test"),
		NewComment(" do something"),
	).AddStatements(NewRawStatement(`fmt.Printf("test\n")`))

	{
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := `case "foo":
	// XXX test test
	// do something
	fmt.Printf("test\n")
`
		assert.Equal(t, expected, gen)
	}

	{
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		expected := `		case "foo":
			// XXX test test
			// do something
			fmt.Printf("test\n")
`
		assert.Equal(t, expected, gen)
	}

	{
		generator = generator.Statements(NewComment("modified"))
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := `case "foo":
	//modified
`
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateCaseRaisesErrorWhenConditionIsEmpty(t *testing.T) {
	generator := NewCase("")
	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.CaseConditionIsEmptyError("").Error(), " ")[0],
	), err.Error())
}

func TestShouldGenerateCaseRaisesErrorWhenStatementsRaisesError(t *testing.T) {
	generator := NewCase(
		`"foo"`,
		NewFunc(nil, NewFuncSignature("")),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
