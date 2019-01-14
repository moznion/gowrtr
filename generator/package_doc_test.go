package generator

import (
	"fmt"
	"log"
)

func ExamplePackage_Generate() {
	generator := NewPackage("mypkg")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
