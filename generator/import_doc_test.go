package generator

import (
	"fmt"
	"log"
)

func ExampleImport_Generate() {
	generator := NewImport("fmt", "os").
		AddImports("exec", "math")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
