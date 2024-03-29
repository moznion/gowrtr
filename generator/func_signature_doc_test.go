package generator

import (
	"fmt"
	"log"
)

func ExampleFuncSignature_Generate() {
	generator := NewFuncSignature(
		"myFunc",
	).TypeParameters(TypeParameters{
		NewTypeParameter("T", "string"),
	}).AddParameters(
		NewFuncParameter("foo", "T"),
		NewFuncParameter("bar", "int"),
	).AddReturnTypeStatements(
		NewFuncReturnTypeWithGenerics("MyStruct", TypeParameterNames{"T", "U"}),
		NewFuncReturnType("error"),
	)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
