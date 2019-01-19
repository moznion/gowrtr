package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateIfCode(t *testing.T) {
	generator := NewIf("i > 0",
		NewComment(" do something"),
		NewRawStatement(`fmt.Printf("%d", i)`),
	)

	{
		expected := `if i > 0 {
	// do something
	fmt.Printf("%d", i)
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		expected := `		if i > 0 {
			// do something
			fmt.Printf("%d", i)
		}
`
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		generator = generator.Statements(NewComment("modified"))
		expected := `if i > 0 {
	//modified
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateIfCodeWithExpandingMethod(t *testing.T) {
	generator := NewIf("i > 0").
		AddStatements(
			NewComment(" XXX: test test"),
			NewComment(" do something"),
		).
		AddStatements(NewRawStatement(`fmt.Printf("%d", i)`))

	expected := `if i > 0 {
	// XXX: test test
	// do something
	fmt.Printf("%d", i)
}
`
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateIfCodeGiveUpWhenStatementRaisesError(t *testing.T) {
	generator := NewIf(
		"i > 0",
		NewFunc(nil, NewFuncSignature("")),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}

func TestShouldGenerateIfAndElseIfAndElseCode(t *testing.T) {
	generator := NewIf("i == 0",
		NewComment(" if"),
	).AddElseIf(
		NewElseIf("i < 0", NewComment(" else if 1")),
		nil,
		NewElseIf("i > 0", NewComment(" else if 2")),
	).Else(NewElse(
		NewComment(" else"),
	))

	{
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		expected := `if i == 0 {
	// if
} else if i < 0 {
	// else if 1
} else if i > 0 {
	// else if 2
} else {
	// else
}
`
		assert.Equal(t, expected, gen)
	}

	{
		gen, err := generator.Generate(2)
		assert.NoError(t, err)
		expected := `		if i == 0 {
			// if
		} else if i < 0 {
			// else if 1
		} else if i > 0 {
			// else if 2
		} else {
			// else
		}
`
		assert.Equal(t, expected, gen)
	}

	{
		generator = generator.ElseIf(
			NewElseIf("ii == 0").Statements(NewComment(" modified")),
		)
		gen, err := generator.Generate(0)

		assert.NoError(t, err)
		expected := `if i == 0 {
	// if
} else if ii == 0 {
	// modified
} else {
	// else
}
`
		assert.Equal(t, expected, gen)
	}

}

func TestShouldGenerateIfRaisesError(t *testing.T) {
	_, err := NewIf("").Generate(0)
	assert.EqualError(t, err, errmsg.IfConditionIsEmptyError().Error())
}

func TestShouldGenerateIfElseIfRaisesError(t *testing.T) {
	generator := NewIf("i == 0",
		NewComment(" if"),
	).AddElseIf(
		NewElseIf("i < 0", NewFuncSignature("")),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}

func TestShouldGenerateIfElseRaisesError(t *testing.T) {
	generator := NewIf("i == 0",
		NewComment(" if"),
	).Else(
		NewElse(NewFuncSignature("")),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}
