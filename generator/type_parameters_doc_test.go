package generator

import (
	"fmt"
	"log"
)

func ExampleTypeArguments_Generate() {
	generated, err := TypeArguments{"int", "string"}.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("// example: f(T, U)")
	fmt.Printf("f%s(intVar, strVar)\n", generated)
}
