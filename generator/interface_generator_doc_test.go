package generator

import (
	"fmt"
	"log"
)

func ExampleInterfaceGenerator_Generate() {
	generator := NewInterfaceGenerator(
		"MyInterface",
		NewFuncSignatureGenerator("fooFunc").
			AddFuncParameters(NewFuncParameter("foo", "string")).
			AddReturnTypes("string", "error"),
	).AddFuncSignatures(
		NewFuncSignatureGenerator("barFunc"),
	)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
