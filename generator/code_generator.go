package gowrtr

import (
	"bytes"
	"io"
	"os/exec"
)

type CodeGeneratable interface {
	Generate(indentLevel int) (string, error)
}

type CodeGenerator struct {
	Statements   []CodeGeneratable
	Gofmt        bool
	GofmtOptions []string
	Goimports    bool
}

func NewCodeGenerator(statements ...CodeGeneratable) *CodeGenerator {
	return &CodeGenerator{
		Statements: statements,
	}
}

func (g *CodeGenerator) AddStatements(statements ...CodeGeneratable) *CodeGenerator {
	return &CodeGenerator{
		Statements:   append(g.Statements, statements...),
		Gofmt:        g.Gofmt,
		GofmtOptions: g.GofmtOptions,
		Goimports:    g.Goimports,
	}
}

func (g *CodeGenerator) EnableGofmt(gofmtOptions ...string) *CodeGenerator {
	return &CodeGenerator{
		Statements:   g.Statements,
		Gofmt:        true,
		GofmtOptions: gofmtOptions,
		Goimports:    g.Goimports,
	}
}

func (g *CodeGenerator) EnableGoimports() *CodeGenerator {
	return &CodeGenerator{
		Statements:   g.Statements,
		Gofmt:        g.Gofmt,
		GofmtOptions: g.GofmtOptions,
		Goimports:    true,
	}
}

func (g *CodeGenerator) Generate(indentLevel int) (string, error) {
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

func (g *CodeGenerator) applyGofmt(generatedCode string) string {
	return applyCodeFormatter(generatedCode, "gofmt", g.GofmtOptions...)
}

func (g *CodeGenerator) applyGoimports(generatedCode string) string {
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
