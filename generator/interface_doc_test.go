package generator

import (
	"fmt"
	"log"
)

func ExampleInterface_Generate() {
	generator := NewInterface(
		"MyInterface",
		NewFuncSignature("fooFunc").
			AddParameters(NewFuncParameter("foo", "T")).
			AddReturnTypes("string", "error"),
	).AddSignatures(
		NewFuncSignature("barFunc"),
	).TypeParameters(TypeParameters{
		NewTypeParameter("T", "any"),
	})

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
