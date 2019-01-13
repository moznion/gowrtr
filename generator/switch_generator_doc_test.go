package generator

import (
	"fmt"
	"log"
)

func ExampleSwitchGenerator_Generate() {
	generator := NewSwitchGenerator("str")
	generator = generator.AddCaseStatements(
		NewCaseGenerator(`"foo"`, NewRawStatementGenerator(`fmt.Printf("str is foo\n")`)),
		NewCaseGenerator(`"bar"`, NewRawStatementGenerator(`fmt.Printf("str is bar\n")`)),
	)
	generator = generator.SetDefaultStatement(
		NewDefaultCaseGenerator(NewRawStatementGenerator(`fmt.Printf("here is default\n")`)),
	)

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
