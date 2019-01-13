package generator

import (
	"fmt"
	"log"
)

func ExampleDefaultCaseGenerator_Generate() {
	generator := NewDefaultCaseGenerator(
		NewCommentGenerator(" XXX test test"),
		NewCommentGenerator(" do something"),
	).AddStatements(NewRawStatementGenerator(`fmt.Printf("test\n")`, true))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
