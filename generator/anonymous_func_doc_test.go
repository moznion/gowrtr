package generator

import (
	"fmt"
	"log"
)

func ExampleAnonymousFunc_Generate() {
	generator := NewAnonymousFunc(
		true,
		NewAnonymousFuncSignature().
			AddParameters(
				NewFuncParameter("foo", "string"),
				NewFuncParameter("bar", "int64"),
			).
			AddReturnTypes("string", "error"),
		NewComment(" do something"),
		NewRawStatement(`fmt.Printf("%d", i)`),
	).Invocation(NewFuncInvocation("foo", "bar"))

	generated, err := generator.Generate(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generated)
}
