gowrtr [![CircleCI](https://circleci.com/gh/moznion/gowrtr.svg?style=svg)](https://circleci.com/gh/moznion/gowrtr) [![codecov](https://codecov.io/gh/moznion/gowrtr/branch/master/graph/badge.svg)](https://codecov.io/gh/moznion/gowrtr) [![GoDoc](https://godoc.org/github.com/moznion/gowrtr/generator?status.svg)](https://godoc.org/github.com/moznion/gowrtr/generator) [![Go Report Card](https://goreportcard.com/badge/github.com/moznion/gowrtr)](https://goreportcard.com/report/github.com/moznion/gowrtr)
==

gowrtr (pronunciation:`go writer`) is a library that supports golang code generation. This library is [go generics](https://go.dev/doc/tutorial/generics) ready!

This library is inspired by [square/javapoet](https://github.com/square/javapoet).

Synopsis
--

Here is a simple example:

```go
package main

import (
	"fmt"

	"github.com/moznion/gowrtr/generator"
)

func main() {
	generator := generator.NewRoot(
		generator.NewComment(" THIS CODE WAS AUTO GENERATED"),
		generator.NewPackage("main"),
		generator.NewNewline(),
	).AddStatements(
		generator.NewFunc(
			nil,
			generator.NewFuncSignature("main"),
		).AddStatements(
			generator.NewRawStatement(`fmt.Println("hello, world!")`),
		),
	).
	Gofmt("-s").
	Goimports()

	generated, err := generator.Generate(0)
	if err != nil {
		panic(err)
	}
	fmt.Println(generated)
}
```

then it generates the golang code like so:

```go
// THIS CODE WAS AUTO GENERATED
package main

import "fmt"

func main() {
        fmt.Println("hello, world!")
}
```

Generics example is here:

```go
package main

import (
	"fmt"

	"github.com/moznion/gowrtr/generator"
)

func main() {
	generated, err := generator.NewRoot(
		generator.NewComment(" THIS CODE WAS AUTO GENERATED"),
		generator.NewPackage("main"),
		generator.NewNewline(),
		generator.NewStruct("MyStruct").
			TypeParameters(TypeParameters{
				generator.NewTypeParameter("T", "any"),
			}).AddField("anything", "T"),
		generator.NewNewline(),
		generator.NewFunc(
			nil,
			generator.NewFuncSignature("foo").
				TypeParameters(TypeParameters{
					generator.NewTypeParameter("U", "any"),
				}).
				ReturnTypeStatements(generator.NewFuncReturnTypeWithTypeParam("MyStruct", []string{"U"})),
		).AddStatements(generator.NewComment("do something")),
		generator.NewNewline(),
		generator.NewFunc(
			generator.NewFuncReceiver("s", "MyStruct", "T"),
			generator.NewFuncSignature("bar").
				ReturnTypeStatements(generator.NewFuncReturnTypeWithTypeParam("MyStruct", []string{"T"})),
		).AddStatements(generator.NewComment("do something")),
	).
	Gofmt("-s").
	Goimports().
	Generate(0)

	if err != nil {
		panic(err)
	}
	fmt.Println(generated)
}
```

then it generates the following golang code:

```go
// THIS CODE WAS AUTO GENERATED
package main

type MyStruct[T any] struct {
	anything T
}

func foo[U any]() MyStruct[U] {
	//do something
}

func (s MyStruct[T]) bar() MyStruct[T] {
	//do something
}
```

And [GoDoc](https://godoc.org/github.com/moznion/gowrtr/generator) shows you a greater number of examples.

Description
--

Please refer to the godoc: [![GoDoc](https://godoc.org/github.com/moznion/gowrtr/generator?status.svg)](https://godoc.org/github.com/moznion/gowrtr/generator)

### Root

- `Root` is an entry point to generate the go code.
- `Root` supports following code formatting on code generating phase. It applies such formatters to generated code.
  - `gofmt`: with `Gofmt(gofmtOptions ...string)`
  - `goimports`: with `Goimports()`

### Immutability

Methods of this library act as immutable. It means it doesn't change any internal state implicitly, so you can take a snapshot of the code generator. That is useful to reuse and derive the code generator instance.

### Debug friendly

This library shows "where is a cause of the error" when code generator raises an error. This means each error message contains a pointer for the error source (i.e. file name and the line number).  This should be helpful for debugging.

Error messages example:

```
[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at /tmp/main.go:22)
```

### Supported syntax

- [x] `package`
- [x] `import`
- [x] `struct`
  - [x] generics type parameters
- [x] `interface`
  - [x] generics type parameters
- [x] [composite literal](https://golang.org/doc/effective_go.html#composite_literals)
- [x] `if`
  - [x] `else if`
  - [x] `else`
- [x] `switch`
  - [x] `case`
  - [x] `default`
- [x] `for`
- [x] code block
- [x] `func`
  - [x] generics type parameters on the signature
  - [x] generics type names on the receiver
  - [x] generics type names on the return types
  - [x] generics types on the invocation
- [x] anonymous func
  - [x] immediately invoking
- one line statement
  - [x] raw
  - [x] newline
  - [x] `return`
  - [x] `comment`

For developers of this library
--

### Setup development environment

```
$ make bootstrap
```

### How to define and generate error messages

Please edit `internal/errmsg/errmsg.go` and execute `make errgen`.

See also: [moznion/go-errgen](https://github.com/moznion/go-errgen)

Blog posts
--

- English: [gowrtr - a library that supports golang code generation](https://moznion.hatenablog.jp/entry/2019/01/15/094236)
- Japanese: [gowrtr - goコード生成支援ライブラリ](https://moznion.hatenadiary.com/entry/2019/01/14/111719)

License
--

```
The MIT License (MIT)
Copyright © 2019 moznion, http://moznion.net/ <moznion@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

