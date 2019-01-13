package generator

import (
	"fmt"
	"log"
)

func ExampleReturnStatementGenerator_Generate() {
	generator := NewReturnStatementGenerator("foo")
	generator = generator.AddReturnItems("err")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
