package generator

import (
	"fmt"
	"log"
)

func ExampleFuncGenerator_Generate() {
	generator := NewFuncGenerator(
		NewFuncReceiverGenerator("m", "*MyStruct"),
		NewFuncSignatureGenerator("myFunc").
			AddFuncParameters(
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			).
			AddReturnTypes("string", "error"),
	).AddStatements(
		NewCommentGenerator(" do something"),
		NewNewlineGenerator(),
		NewReturnStatementGenerator("foo+bar", "nil"),
	)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
