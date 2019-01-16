package generator

import (
	"fmt"
	"log"
)

func ExampleFor_Generate() {
	generator := NewFor(
		"i := 0; i < foo; i++",
		NewComment(" do something"),
	).AddStatements(NewRawStatement(`fmt.Printf("%d", i)`))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
