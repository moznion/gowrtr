package generator

import (
	"fmt"
	"log"
)

func ExampleFuncReceiver_Generate() {
	funcReceiver := NewFuncReceiver("f", "*Foo")

	generated, err := funcReceiver.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
