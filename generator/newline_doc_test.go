package generator

import (
	"fmt"
	"log"
)

func ExampleNewline_Generate() {
	generator := NewNewline()

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
