package generator

import (
	"fmt"
	"log"
)

func ExampleCase_Generate() {
	generator := NewCase(`"foo"`, NewComment(" this is foo")).
		AddStatements(NewRawStatement(`fmt.Printf("this is foo\n")`))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(generated)
}
