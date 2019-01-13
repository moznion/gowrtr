package generator

import (
	"fmt"
	"log"
)

func ExampleForGenerator_Generate() {
	generator := NewForGenerator(
		"i := 0; i < foo; i++",
		NewCommentGenerator(" do something"),
	).AddStatements(NewRawStatementGenerator(`fmt.Printf("%d", i)`, true))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
