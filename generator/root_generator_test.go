package gowrtr

import (
	"regexp"
	"testing"

	"github.com/moznion/gowrtr/internal/errmsg"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateCode(t *testing.T) {
	expected := `// THIS CODE WAS AUTO GENERATED

package mypkg

import (
	"fmt"
)

type MyInterface interface {
	MyFunc(foo string) (string, error)
}

type MyStruct struct {
	Foo string
	Bar int64
}

func (m *MyStruct) MyFunc(foo string) (string, error) {
	{
		str := 		func(bar string) string {
			return bar
		}(foo)

		if str == "" {
			for i := 0; i < 3; i++ {
				fmt.Printf("%d\n", i)
			}
		}
		return str, nil
	}
}
`
	myFuncSignature := NewFuncSignatureGenerator("MyFunc").
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewRootGenerator(
		NewCommentGenerator(" THIS CODE WAS AUTO GENERATED"),
		NewNewlineGenerator(),
		NewPackageGenerator("mypkg"),
		NewNewlineGenerator(),
		NewImportGenerator("fmt"),
		NewNewlineGenerator(),
		NewInterfaceGenerator("MyInterface").
			AddFuncSignature(myFuncSignature),
		NewNewlineGenerator(),
		NewStructGenerator("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewlineGenerator(),
	).AddStatements(
		NewFuncGenerator(
			NewFuncReceiverGenerator("m", "*MyStruct"),
			NewFuncSignatureGenerator("MyFunc").
				AddFuncParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlockGenerator(
				NewRawStatementGenerator("str := "),
				NewInlineFuncGenerator(
					false,
					NewInlineFuncSignatureGenerator().
						AddFuncParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatementGenerator("bar"),
				).AddFuncInvocation(NewFuncInvocationGenerator("foo")),
				NewNewlineGenerator(),
				NewIfGenerator(`str == ""`).
					AddStatements(
						NewForGenerator(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatementGenerator(`fmt.Printf("%d\n", i)`, true),
						),
					),
				NewReturnStatementGenerator("str", "nil"),
			),
		),
	)

	generated, err := generator.EnableSyntaxChecking().Generate(0)

	assert.NoError(t, err)
	assert.Equal(t, expected, generated)
}

func TestShouldGenerateCodeWithIndent(t *testing.T) {
	expected := `		// THIS CODE WAS AUTO GENERATED

		package mypkg

		import (
			"fmt"
		)

		type MyInterface interface {
			MyFunc(foo string) (string, error)
		}

		type MyStruct struct {
			Foo string
			Bar int64
		}

		func (m *MyStruct) MyFunc(foo string) (string, error) {
			{
				str := 				func(bar string) string {
					return bar
				}(foo)

				if str == "" {
					for i := 0; i < 3; i++ {
						fmt.Printf("%d\n", i)
					}
				}
				return str, nil
			}
		}
`
	myFuncSignature := NewFuncSignatureGenerator("MyFunc").
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewRootGenerator(
		NewCommentGenerator(" THIS CODE WAS AUTO GENERATED"),
		NewNewlineGenerator(),
		NewPackageGenerator("mypkg"),
		NewNewlineGenerator(),
		NewImportGenerator("fmt"),
		NewNewlineGenerator(),
		NewInterfaceGenerator("MyInterface").
			AddFuncSignature(myFuncSignature),
		NewNewlineGenerator(),
		NewStructGenerator("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewlineGenerator(),
	).AddStatements(
		NewFuncGenerator(
			NewFuncReceiverGenerator("m", "*MyStruct"),
			NewFuncSignatureGenerator("MyFunc").
				AddFuncParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlockGenerator(
				NewRawStatementGenerator("str := "),
				NewInlineFuncGenerator(
					false,
					NewInlineFuncSignatureGenerator().
						AddFuncParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatementGenerator("bar"),
				).AddFuncInvocation(NewFuncInvocationGenerator("foo")),
				NewNewlineGenerator(),
				NewIfGenerator(`str == ""`).
					AddStatements(
						NewForGenerator(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatementGenerator(`fmt.Printf("%d\n", i)`, true),
						),
					),
				NewReturnStatementGenerator("str", "nil"),
			),
		),
	)

	generated, err := generator.Generate(2)

	assert.NoError(t, err)
	assert.Equal(t, expected, generated)
}

func TestShouldGenerateCodeWithGofmt(t *testing.T) {
	expected := `// THIS CODE WAS AUTO GENERATED

package mypkg

import (
	"fmt"
)

type MyInterface interface {
	MyFunc(foo string) (string, error)
}

type MyStruct struct {
	Foo string
	Bar int64
}

func (m *MyStruct) MyFunc(foo string) (string, error) {
	{
		str := func(bar string) string {
			return bar
		}(foo)

		if str == "" {
			for i := 0; i < 3; i++ {
				fmt.Printf("%d\n", i)
			}
		}
		return str, nil
	}
}
`
	myFuncSignature := NewFuncSignatureGenerator("MyFunc").
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewRootGenerator(
		NewCommentGenerator(" THIS CODE WAS AUTO GENERATED"),
		NewNewlineGenerator(),
		NewPackageGenerator("mypkg"),
		NewNewlineGenerator(),
		NewImportGenerator("fmt"),
		NewNewlineGenerator(),
		NewInterfaceGenerator("MyInterface").
			AddFuncSignature(myFuncSignature),
		NewNewlineGenerator(),
		NewStructGenerator("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewlineGenerator(),
		NewFuncGenerator(
			NewFuncReceiverGenerator("m", "*MyStruct"),
			NewFuncSignatureGenerator("MyFunc").
				AddFuncParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlockGenerator(
				NewRawStatementGenerator("str := "),
				NewInlineFuncGenerator(
					false,
					NewInlineFuncSignatureGenerator().
						AddFuncParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatementGenerator("bar"),
				).AddFuncInvocation(NewFuncInvocationGenerator("foo")),
				NewNewlineGenerator(),
				NewIfGenerator(`str == ""`).
					AddStatements(
						NewForGenerator(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatementGenerator(`fmt.Printf("%d\n", i)`, true),
						),
					),
				NewReturnStatementGenerator("str", "nil"),
			),
		),
	).EnableGofmt("-s")

	generated, err := generator.Generate(0)

	assert.NoError(t, err)
	assert.Equal(t, expected, generated)
}

func TestShouldGenerateCodeWithGoimport(t *testing.T) {
	expected := `// THIS CODE WAS AUTO GENERATED

package mypkg

import "fmt"

type MyInterface interface {
	MyFunc(foo string) (string, error)
}

type MyStruct struct {
	Foo string
	Bar int64
}

func (m *MyStruct) MyFunc(foo string) (string, error) {
	{
		str := func(bar string) string {
			return bar
		}(foo)

		if str == "" {
			for i := 0; i < 3; i++ {
				fmt.Printf("%d\n", i)
			}
		}
		return str, nil
	}
}
`
	myFuncSignature := NewFuncSignatureGenerator("MyFunc").
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewRootGenerator(
		NewCommentGenerator(" THIS CODE WAS AUTO GENERATED"),
		NewNewlineGenerator(),
		NewPackageGenerator("mypkg"),
		NewInterfaceGenerator("MyInterface").
			AddFuncSignature(myFuncSignature),
		NewNewlineGenerator(),
		NewStructGenerator("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewlineGenerator(),
		NewFuncGenerator(
			NewFuncReceiverGenerator("m", "*MyStruct"),
			NewFuncSignatureGenerator("MyFunc").
				AddFuncParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlockGenerator(
				NewRawStatementGenerator("str := "),
				NewInlineFuncGenerator(
					false,
					NewInlineFuncSignatureGenerator().
						AddFuncParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatementGenerator("bar"),
				).AddFuncInvocation(NewFuncInvocationGenerator("foo")),
				NewNewlineGenerator(),
				NewIfGenerator(`str == ""`).
					AddStatements(
						NewForGenerator(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatementGenerator(`fmt.Printf("%d\n", i)`, true),
						),
					),
				NewReturnStatementGenerator("str", "nil"),
			),
		),
	).EnableGoimports()

	generated, err := generator.Generate(0)

	assert.NoError(t, err)
	assert.Equal(t, expected, generated)
}

func TestShouldGenerateCodeRaisesError(t *testing.T) {
	generator := NewRootGenerator(
		NewCommentGenerator(" THIS CODE WAS AUTO GENERATED"),
		NewNewlineGenerator(),
		NewPackageGenerator("mypkg"),
		NewNewlineGenerator(),
		NewFuncGenerator(
			nil,
			nil,
		),
	)

	_, err := generator.Generate(0)
	assert.EqualError(t, err, errmsg.FuncSignatureIsNilError().Error())
}

func TestShouldGenerateCodeRaiseErrorWhenCodeFormatterIsExited(t *testing.T) {
	{
		generator := NewRootGenerator(
			NewRawStatementGenerator("\timport something", true),
		).EnableSyntaxChecking()

		_, err := generator.Generate(0)
		assert.Regexp(t, regexp.MustCompile(`^\[GOWRTR-13\] code formatter raises error: command="gofmt -e".+`), err.Error())
	}

	{
		generator := NewRootGenerator(
			NewRawStatementGenerator("\timport something", true),
		).EnableGofmt()

		_, err := generator.Generate(0)
		assert.Regexp(t, regexp.MustCompile(`^\[GOWRTR-13\] code formatter raises error: command="gofmt".+`), err.Error())
	}

	{
		generator := NewRootGenerator(
			NewRawStatementGenerator("\timport something", true),
		).EnableGoimports()

		_, err := generator.Generate(0)
		assert.Regexp(t, regexp.MustCompile(`^\[GOWRTR-13\] code formatter raises error: command="goimports".+`), err.Error())
	}
}
