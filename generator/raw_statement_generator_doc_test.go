package generator

import (
	"fmt"
	"log"
)

func ExampleRawStatementGenerator_Generate() {
	generator := NewRawStatementGenerator("i := 1 + 1", true)
	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
