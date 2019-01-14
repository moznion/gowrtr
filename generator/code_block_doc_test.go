package generator

import (
	"fmt"
	"log"
)

func ExampleCodeBlock_Generate() {
	generator := NewCodeBlock(NewComment(" do something"))
	generator = generator.AddStatements(NewRawStatement(`fmt.Printf("code block\n")`))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
