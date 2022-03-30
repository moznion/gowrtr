package generator

import (
	"fmt"
	"log"
)

func ExampleStruct_Generate() {
	generator := NewStruct("MyStruct")
	generator = generator.
		TypeParameters(TypeParameters{NewTypeParameter("T", "string")}).
		AddField("foo", "T").
		AddField("bar", "int64", `custom:"tag"`)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
