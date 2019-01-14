package generator

import (
	"fmt"
	"log"
)

func ExampleIf_Generate() {
	generator := NewIf("i == 0",
		NewComment(" if"),
	).AddElseIfBlocks(
		NewElseIf("i < 0", NewComment(" else if 1")),
		nil,
		NewElseIf("i > 0", NewComment(" else if 2")),
	).SetElseBlock(NewElse(
		NewComment(" else"),
	))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
