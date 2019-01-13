package generator

import (
	"fmt"
	"log"
)

func ExampleFuncReceiverGenerator_Generate() {
	funcReceiver := NewFuncReceiverGenerator("f", "*Foo")

	generated, err := funcReceiver.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
