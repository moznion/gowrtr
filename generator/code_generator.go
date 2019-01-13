package gowrtr

import (
	"bytes"
	"io"
	"os/exec"
)

type StatementGenerator interface {
	Generate(indentLevel int) (string, error)
}

type RootGenerator struct {
	Statements   []StatementGenerator
	Gofmt        bool
	GofmtOptions []string
	Goimports    bool
}

func NewCodeGenerator(statements ...StatementGenerator) *RootGenerator {
	return &RootGenerator{
		Statements: statements,
	}
}

func (g *RootGenerator) AddStatements(statements ...StatementGenerator) *RootGenerator {
	return &RootGenerator{
		Statements:   append(g.Statements, statements...),
		Gofmt:        g.Gofmt,
		GofmtOptions: g.GofmtOptions,
		Goimports:    g.Goimports,
	}
}

func (g *RootGenerator) EnableGofmt(gofmtOptions ...string) *RootGenerator {
	return &RootGenerator{
		Statements:   g.Statements,
		Gofmt:        true,
		GofmtOptions: gofmtOptions,
		Goimports:    g.Goimports,
	}
}

func (g *RootGenerator) EnableGoimports() *RootGenerator {
	return &RootGenerator{
		Statements:   g.Statements,
		Gofmt:        g.Gofmt,
		GofmtOptions: g.GofmtOptions,
		Goimports:    true,
	}
}

func (g *RootGenerator) Generate(indentLevel int) (string, error) {
	generatedCode := ""

	for _, statement := range g.Statements {
		gen, err := statement.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		generatedCode += gen
	}

	if g.Gofmt {
		generatedCode = g.applyGofmt(generatedCode)
	}

	if g.Goimports {
		generatedCode = g.applyGoimports(generatedCode)
	}

	return generatedCode, nil
}

func (g *RootGenerator) applyGofmt(generatedCode string) string {
	return applyCodeFormatter(generatedCode, "gofmt", g.GofmtOptions...)
}

func (g *RootGenerator) applyGoimports(generatedCode string) string {
	return applyCodeFormatter(generatedCode, "goimports")
}

func applyCodeFormatter(generatedCode string, formatterCmdName string, formatterOpts ...string) string {
	echoCmd := exec.Command("echo", generatedCode)
	formatterCmd := exec.Command(formatterCmdName, formatterOpts...)

	r, w := io.Pipe()
	echoCmd.Stdout = w
	formatterCmd.Stdin = r

	var out bytes.Buffer
	formatterCmd.Stdout = &out

	echoCmd.Start()
	formatterCmd.Start()
	echoCmd.Wait()
	w.Close()
	formatterCmd.Wait()

	return out.String()
}

func buildIndent(indentLevel int) string {
	indent := ""
	for i := 0; i < indentLevel; i++ {
		indent += "\t"
	}
	return indent
}
