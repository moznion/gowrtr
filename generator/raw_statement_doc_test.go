package generator

import (
	"fmt"
	"log"
)

func ExampleRawStatement_Generate() {
	generator := NewRawStatement("i := 1 + 1")
	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
