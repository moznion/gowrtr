package generator

import (
	"fmt"
	"log"
)

func ExampleReturnStatement_Generate() {
	generator := NewReturnStatement("foo")
	generator = generator.AddReturnItems("err")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
