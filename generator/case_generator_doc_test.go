package generator

import (
	"fmt"
	"log"
)

func ExampleCaseGenerator_Generate() {
	generator := NewCaseGenerator(`"foo"`, NewCommentGenerator(" this is foo")).
		AddStatements(NewRawStatementGenerator(`fmt.Printf("this is foo\n")`))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(generated)
}
