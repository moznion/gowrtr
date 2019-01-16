package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateDefaultCase(t *testing.T) {
	generator := NewDefaultCase(
		NewComment(" XXX test test"),
		NewComment(" do something"),
	).AddStatements(NewRawStatement(`fmt.Printf("test\n")`))

	{
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := `default:
	// XXX test test
	// do something
	fmt.Printf("test\n")
`
		assert.Equal(t, expected, gen)
	}

	{
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		expected := `		default:
			// XXX test test
			// do something
			fmt.Printf("test\n")
`
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateDefaultCaseRaisesErrorWhenStatementsRaisesError(t *testing.T) {
	generator := NewDefaultCase(
		NewFunc(nil, NewFuncSignature("")),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
