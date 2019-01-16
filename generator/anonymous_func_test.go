package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateAnonymousFunc(t *testing.T) {
	generator := NewAnonymousFunc(
		false,
		NewAnonymousFuncSignature(),
	)

	{
		expected := `func() {
}
`
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}

	{
		expected := `func() {
	// do something
	fmt.Printf("%d", i)
}
`
		generator = generator.AddStatements(
			NewComment(" do something"),
			NewRawStatement(`fmt.Printf("%d", i)`),
		)
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateAnonymousFuncWithSetterMethod(t *testing.T) {

	expected := `func() {
	// do something
	fmt.Printf("%d", i)
}
`
	generator := NewAnonymousFunc(false,
		NewAnonymousFuncSignature(),
		NewComment(" do something"),
		NewRawStatement(`fmt.Printf("%d", i)`),
	)
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)

	generator = generator.Statements(NewComment("modified"))
	expected = `func() {
	//modified
}
`
	gen, err = generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateAnonymousFuncWithSignature(t *testing.T) {
	generator := NewAnonymousFunc(
		false,
		NewAnonymousFuncSignature().
			AddFuncParameters(
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int64"),
			).
			AddReturnTypes("string", "error"),
		NewComment(" do something"),
		NewRawStatement(`fmt.Printf("%d", i)`),
	)

	expected := `func(foo string, bar int64) (string, error) {
	// do something
	fmt.Printf("%d", i)
}
`
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateAnonymousGoFuncWithInvocation(t *testing.T) {
	generator := NewAnonymousFunc(
		true,
		NewAnonymousFuncSignature().
			AddFuncParameters(
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int64"),
			).
			AddReturnTypes("string", "error"),
		NewComment(" do something"),
		NewRawStatement(`fmt.Printf("%d", i)`),
	).SetFuncInvocation(NewFuncInvocation("foo", "bar"))

	expected := `go func(foo string, bar int64) (string, error) {
	// do something
	fmt.Printf("%d", i)
}(foo, bar)
`
	gen, err := generator.Generate(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, gen)
}

func TestShouldGenerateAnonymousFuncRaisesErrorWhenAnonymousFuncSignatureIsNil(t *testing.T) {
	generator := NewAnonymousFunc(
		false,
		nil,
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.AnonymousFuncSignatureIsNilError().Error())
}

func TestShouldGenerateAnonymousFuncRaisesErrorWhenAnonymousFuncSignatureRaisesError(t *testing.T) {
	generator := NewAnonymousFunc(
		false,
		NewAnonymousFuncSignature().AddFuncParameters(
			NewFuncParameter("", "string"),
		),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncParameterNameIsEmptyErr().Error())
}

func TestShouldGenerateAnonymousFuncRaisesErrorWhenStatementRaisesError(t *testing.T) {
	generator := NewAnonymousFunc(
		false,
		NewAnonymousFuncSignature(),
		NewFunc(nil, NewFuncSignature("")),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}

func TestShouldGenerateAnonymousFuncRaisesErrorWhenFuncInvocationRaisesError(t *testing.T) {
	generator := NewAnonymousFunc(
		false,
		NewAnonymousFuncSignature(),
	).SetFuncInvocation(NewFuncInvocation(""))
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncInvocationParameterIsEmptyError().Error())
}
