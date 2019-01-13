package generator

import (
	"fmt"
	"log"
)

func ExamplePackageGenerator_Generate() {
	generator := NewPackageGenerator("mypkg")

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
