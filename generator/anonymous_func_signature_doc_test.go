package generator

import (
	"fmt"
	"log"
)

func ExampleAnonymousFuncSignature_Generate() {
	generator := NewAnonymousFuncSignature().
		AddParameters(
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
