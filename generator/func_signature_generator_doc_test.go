package generator

import (
	"fmt"
	"log"
)

func ExampleFuncSignatureGenerator_Generate() {
	generator := NewFuncSignatureGenerator(
		"myFunc",
	).AddFuncParameters(
		NewFuncParameter("foo", "string"),
		NewFuncParameter("bar", "int"),
	).AddReturnTypes("string", "error")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
