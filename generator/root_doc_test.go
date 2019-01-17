package generator

import (
	"fmt"
	"log"
)

func ExampleRoot_Generate() {
	myFuncSignature := NewFuncSignature("MyFunc").
		AddParameters(
			NewFuncParameter("foo", "string"),
		).
		AddReturnTypes("string", "error")

	generator := NewRoot(
		NewComment(" THIS CODE WAS AUTO GENERATED"),
		NewNewline(),
		NewPackage("mypkg"),
		NewInterface("MyInterface").
			AddSignatures(myFuncSignature),
		NewNewline(),
		NewStruct("MyStruct").
			AddField("Foo", "string").
			AddField("Bar", "int64"),
		NewNewline(),
	).AddStatements(
		NewFunc(
			NewFuncReceiver("m", "*MyStruct"),
			NewFuncSignature("MyFunc").
				AddParameters(
					NewFuncParameter("foo", "string"),
				).
				AddReturnTypes("string", "error"),
		).AddStatements(
			NewCodeBlock(
				NewRawStatement("str := ").WithNewline(false),
				NewAnonymousFunc(
					false,
					NewAnonymousFuncSignature().
						AddParameters(NewFuncParameter("bar", "string")).
						AddReturnTypes("string"),
					NewReturnStatement("bar"),
				).Invocation(NewFuncInvocation("foo")),
				NewNewline(),
				NewIf(`str == ""`).
					AddStatements(
						NewFor(`i := 0; i < 3; i++`).AddStatements(
							NewRawStatement(`fmt.Printf("%d\n", i)`),
						),
					),
				NewNewline(),
				NewSwitch("str").
					AddCaseStatements(
						NewCase(
							`""`,
							NewComment(" empty string"),
						),
						NewCase(
							`"foo"`,
							NewComment(" foo string"),
						),
					).
					SetDefaultStatement(
						NewDefaultCase(NewComment(" default")),
					),
				NewNewline(),
				NewReturnStatement("str", "nil"),
			),
		),
	)

	generated, err := generator.
		EnableGofmt("-s").
		EnableGoimports().
		Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
