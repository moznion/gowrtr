package generator

import (
	"fmt"
	"log"
)

func ExampleFuncInvocation_Generate() {
	generator := NewFuncInvocation("foo").AddParameters("bar")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}

func ExampleFuncInvocation_Generate_withGenericsTypes() {
	generator := NewFuncInvocation("foo").AddParameters("bar").GenericsTypes("string", "int64")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
