package generator

import (
	"fmt"
	"log"
)

func ExampleElseIf_Generate() {
	generator := NewElseIf("i > 0").
		AddStatements(
			NewComment(" XXX: test test"),
			NewComment(" do something"),
		).
		AddStatements(NewRawStatement(`fmt.Printf("%d", i)`))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
