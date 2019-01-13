package generator

import (
	"fmt"
	"log"
)

func ExampleCodeBlockGenerator_Generate() {
	generator := NewCodeBlockGenerator(NewCommentGenerator(" do something"))
	generator = generator.AddStatements(NewRawStatementGenerator(`fmt.Printf("code block\n")`))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
