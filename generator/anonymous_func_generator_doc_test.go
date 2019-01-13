package generator

import (
	"fmt"
	"log"
)

func ExampleAnonymousFuncGenerator_Generate() {
	generator := NewAnonymousFuncGenerator(
		true,
		NewAnonymousFuncSignatureGenerator().
			AddFuncParameters(
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int64"),
			).
			AddReturnTypes("string", "error"),
		NewCommentGenerator(" do something"),
		NewRawStatementGenerator(`fmt.Printf("%d", i)`, true),
	).SetFuncInvocation(NewFuncInvocationGenerator("foo", "bar"))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
