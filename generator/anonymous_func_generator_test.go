package generator

import (
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateAnonymousFunc(t *testing.T) {
	generator := NewAnonymousFuncGenerator(
		false,
		NewAnonymousFuncSignatureGenerator(),
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
			NewCommentGenerator(" do something"),
			NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
		)
		gen, err := generator.Generate(0)
		assert.NoError(t, err)
		assert.Equal(t, expected, gen)
	}
}

func TestShouldGenerateAnonymousFuncWithSignature(t *testing.T) {
	generator := NewAnonymousFuncGenerator(
		false,
		NewAnonymousFuncSignatureGenerator().
			AddFuncParameters(
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int64"),
			).
			AddReturnTypes("string", "error"),
		NewCommentGenerator(" do something"),
		NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
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
	generator := NewAnonymousFuncGenerator(
		true,
		NewAnonymousFuncSignatureGenerator().
			AddFuncParameters(
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int64"),
			).
			AddReturnTypes("string", "error"),
		NewCommentGenerator(" do something"),
		NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
	).SetFuncInvocation(NewFuncInvocationGenerator("foo", "bar"))

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
	generator := NewAnonymousFuncGenerator(
		false,
		nil,
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.AnonymousFuncSignatureIsNilError().Error())
}

func TestShouldGenerateAnonymousFuncRaisesErrorWhenAnonymousFuncSignatureGeneratorRaisesError(t *testing.T) {
	generator := NewAnonymousFuncGenerator(
		false,
		NewAnonymousFuncSignatureGenerator().AddFuncParameters(
			NewFuncParameter("", "string"),
		),
	)
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncParameterNameIsEmptyErr().Error())
}

func TestShouldGenerateAnonymousFuncRaisesErrorWhenStatementRaisesError(t *testing.T) {
	generator := NewAnonymousFuncGenerator(
		false,
		NewAnonymousFuncSignatureGenerator(),
		NewFuncGenerator(nil, NewFuncSignatureGenerator("")),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncNameIsEmptyError().Error())
}

func TestShouldGenerateAnonymousFuncRaisesErrorWhenFuncInvocationGeneratorRaisesError(t *testing.T) {
	generator := NewAnonymousFuncGenerator(
		false,
		NewAnonymousFuncSignatureGenerator(),
	).SetFuncInvocation(NewFuncInvocationGenerator(""))
	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncInvocationParameterIsEmptyError().Error())
}
