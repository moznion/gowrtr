package generator

import (
	"fmt"
	"log"
)

func ExampleNewlineGenerator_Generate() {
	generator := NewNewlineGenerator()

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
