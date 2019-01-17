package generator

import (
	"fmt"
	"log"
)

func ExampleFuncSignature_Generate() {
	generator := NewFuncSignature(
		"myFunc",
	).AddParameters(
		NewFuncParameter("foo", "string"),
		NewFuncParameter("bar", "int"),
	).AddReturnTypes("string", "error")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
