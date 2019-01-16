package generator

import (
	"fmt"
	"log"
)

func ExampleElse_Generate() {
	generator := NewElse(
		NewComment(" XXX test test"),
		NewComment(" do something"),
	).AddStatements(
		NewRawStatement(`fmt.Printf("%d", i)`),
	)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
