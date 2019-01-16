package generator

import (
	"fmt"
	"log"
)

func ExampleDefaultCase_Generate() {
	generator := NewDefaultCase(
		NewComment(" XXX test test"),
		NewComment(" do something"),
	).AddStatements(NewRawStatement(`fmt.Printf("test\n")`))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
