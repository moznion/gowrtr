package generator

import (
	"fmt"
	"log"
)

func ExampleImportGenerator_Generate() {
	generator := NewImportGenerator("fmt", "os").
		AddImports("exec", "math")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
