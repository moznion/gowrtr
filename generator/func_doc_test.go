package generator

import (
	"fmt"
	"log"
)

func ExampleFunc_Generate() {
	generator := NewFunc(
		NewFuncReceiver("m", "*MyStruct"),
		NewFuncSignature("myFunc").
			AddParameters(
				NewFuncParameter("foo", ""),
				NewFuncParameter("bar", "string"),
			).
			AddReturnTypes("string", "error"),
	).AddStatements(
		NewComment(" do something"),
		NewNewline(),
		NewReturnStatement("foo+bar", "nil"),
	)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
