package generator

import (
	"fmt"
	"log"
)

func ExampleInterface_Generate() {
	generator := NewInterface(
		"MyInterface",
		NewFuncSignature("fooFunc").
			AddFuncParameters(NewFuncParameter("foo", "string")).
			AddReturnTypes("string", "error"),
	).AddFuncSignatures(
		NewFuncSignature("barFunc"),
	)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
