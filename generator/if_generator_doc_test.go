package generator

import (
	"fmt"
	"log"
)

func ExampleIfGenerator_Generate() {
	generator := NewIfGenerator("i == 0",
		NewCommentGenerator(" if"),
	).AddElseIfBlocks(
		NewElseIfGenerator("i < 0", NewCommentGenerator(" else if 1")),
		nil,
		NewElseIfGenerator("i > 0", NewCommentGenerator(" else if 2")),
	).SetElseBlock(NewElseGenerator(
		NewCommentGenerator(" else"),
	))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
