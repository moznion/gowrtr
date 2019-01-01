package gowrtr

import (
	"fmt"
)

type Import struct {
	Names []string
}

func NewImport(names ...string) *Import {
	return &Import{
		Names: names,
	}
}

func (imp *Import) String() string {
	if len(imp.Names) <= 0 {
		return ""
	}

	stmt := "import (\n"
	for _, name := range imp.Names {
		stmt += fmt.Sprintf("\t\"%s\"\n", name)
	}
	stmt += ")"

	return stmt
}
