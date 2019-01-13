package generator

import (
	"fmt"
	"log"
)

func ExampleCommentGenerator_Generate() {
	generator := NewCommentGenerator("this is one line comment")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
