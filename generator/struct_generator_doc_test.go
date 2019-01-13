package generator

import (
	"fmt"
	"log"
)

func ExampleStructGenerator_Generate() {
	generator := NewStructGenerator("MyStruct")
	generator = generator.
		AddField("foo", "string").
		AddField("bar", "int64", `custom:"tag"`)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
