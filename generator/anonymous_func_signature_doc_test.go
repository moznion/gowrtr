package generator

import (
	"fmt"
	"log"
)

func ExampleAnonymousFuncSignature_Generate() {
	generator := NewAnonymousFuncSignature().
		AddFuncParameters(
			NewFuncParameter("foo", "string"),
			NewFuncParameter("bar", "int64"),
		).
		AddReturnTypes("string", "error")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
