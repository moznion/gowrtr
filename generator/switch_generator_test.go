package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"
	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateSwitch(t *testing.T) {
	generator := NewSwitchGenerator("foo")

	{
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := `switch foo {
}
`
		assert.Equal(t, expected, gen)
	}

	generator = generator.AddCaseStatements(
		NewCaseGenerator("1", NewCommentGenerator(" one")),
		nil,
		NewCaseGenerator("2", NewCommentGenerator(" two")),
	).SetDefaultStatement(
		NewDefaultCaseGenerator(NewCommentGenerator(" default")),
	)

	{
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := `switch foo {
case 1:
	// one
case 2:
	// two
default:
	// default
}
`
		assert.Equal(t, expected, gen)
	}

	{
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		expected := `		switch foo {
		case 1:
			// one
		case 2:
			// two
		default:
			// default
		}
`
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateSwitchRaisesErrorWhenCaseGeneratorRaisesError(t *testing.T) {
	generator := NewSwitchGenerator("foo").AddCaseStatements(
		NewCaseGenerator(""),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.CaseConditionIsEmptyError().Error())
}

func TestShouldGenerateSwitchRaisesErrorWhenDefaultGeneratorRaisesError(t *testing.T) {
	generator := NewSwitchGenerator("foo").SetDefaultStatement(
		NewDefaultCaseGenerator(
			NewFuncGenerator(nil, NewFuncSignatureGenerator("")),
		),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
