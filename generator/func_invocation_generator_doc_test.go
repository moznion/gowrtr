package generator

import (
	"fmt"
	"log"
)

func ExampleFuncInvocationGenerator_Generate() {
	generator := NewFuncInvocationGenerator("foo").AddParameters("bar")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
