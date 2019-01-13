package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCase(t *testing.T) {
	generator := NewCaseGenerator(
		`"foo"`,
		NewCommentGenerator(" XXX test test"),
		NewCommentGenerator(" do something"),
	).AddStatements(NewRawStatementGenerator(`fmt.Printf("test\n")`, true))

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
}

func TestShouldGenerateCaseRaisesErrorWhenConditionIsEmpty(t *testing.T) {
	generator := NewCaseGenerator("")
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.CaseConditionIsEmptyError().Error())
}

func TestShouldGenerateCaseRaisesErrorWhenStatementsRaisesError(t *testing.T) {
	generator := NewCaseGenerator(
		`"foo"`,
		NewFuncGenerator(nil, NewFuncSignatureGenerator("")),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
