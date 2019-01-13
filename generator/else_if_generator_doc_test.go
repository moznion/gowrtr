package generator

import (
	"fmt"
	"log"
)

func ExampleElseIfGenerator_Generate() {
	generator := NewElseIfGenerator("i > 0").
		AddStatements(
			NewCommentGenerator(" XXX: test test"),
			NewCommentGenerator(" do something"),
		).
		AddStatements(NewRawStatementGenerator(`fmt.Printf("%d", i)`, true))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
