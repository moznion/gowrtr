package generator

import (
	"regexp"
	"strings"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateElseCode(t *testing.T) {
	generator := NewElse(
		NewComment(" XXX test test"),
		NewComment(" do something"),
	).AddStatements(
		NewRawStatement(`fmt.Printf("%d", i)`),
	)

	{
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := ` else {
	// XXX test test
	// do something
	fmt.Printf("%d", i)
}`
		assert.Equal(t, expected, gen)
	}

	{
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		expected := ` else {
			// XXX test test
			// do something
			fmt.Printf("%d", i)
		}`
		assert.Equal(t, expected, gen)
	}

	{
		generator = generator.Statements(NewComment("modified"))
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := ` else {
	//modified
}`
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateElseCodeRaisesError(t *testing.T) {
	generator := NewElse(
		NewFunc(nil, NewFuncSignature("")),
	)
	_, err := generator.Generate(0)
	assert.Regexp(t, regexp.MustCompile(
		`^\`+strings.Split(errmsg.FuncNameIsEmptyError("").Error(), " ")[0],
	), err.Error())
}
