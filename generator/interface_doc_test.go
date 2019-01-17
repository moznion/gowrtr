package generator

import (
	"fmt"
	"log"
)

func ExampleInterface_Generate() {
	generator := NewInterface(
		"MyInterface",
		NewFuncSignature("fooFunc").
			AddParameters(NewFuncParameter("foo", "string")).
			AddReturnTypes("string", "error"),
	).AddSignatures(
		NewFuncSignature("barFunc"),
	)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
