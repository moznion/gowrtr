package generator

import (
	"bytes"
	"io"
	"os/exec"
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// Root is a code generator for the entry point.
type Root struct {
	statements     []Statement
	gofmt          bool
	gofmtOptions   []string
	goimports      bool
	syntaxChecking bool
}

// NewRoot generates a new `Root`.
func NewRoot(statements ...Statement) *Root {
	return &Root{
		statements: statements,
	}
}

// AddStatements adds statements to Root. This does *not* set, just add.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) AddStatements(statements ...Statement) *Root {
	return &Root{
		statements:     append(g.statements, statements...),
		gofmt:          g.gofmt,
		gofmtOptions:   g.gofmtOptions,
		goimports:      g.goimports,
		syntaxChecking: g.syntaxChecking,
	}
}

// Statements sets statements to Root. This does *not* add, just set.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) Statements(statements ...Statement) *Root {
	return &Root{
		statements:     statements,
		gofmt:          g.gofmt,
		gofmtOptions:   g.gofmtOptions,
		goimports:      g.goimports,
		syntaxChecking: g.syntaxChecking,
	}
}

// EnableGofmt enables `gofmt`. If `gofmt` is enabled, it applies `gofmt` on code generation phase.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) EnableGofmt(gofmtOptions ...string) *Root {
	return &Root{
		statements:     g.statements,
		gofmt:          true,
		gofmtOptions:   gofmtOptions,
		goimports:      g.goimports,
		syntaxChecking: g.syntaxChecking,
	}
}

// EnableGoimports enables `goimports`. If `goimports` is enabled, it applies `goimports` on code generation phase.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) EnableGoimports() *Root {
	return &Root{
		statements:     g.statements,
		gofmt:          g.gofmt,
		gofmtOptions:   g.gofmtOptions,
		goimports:      true,
		syntaxChecking: g.syntaxChecking,
	}
}

// EnableSyntaxChecking enables syntax checking. If this option is enabled, it checks the syntax of the code on code generation phase.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) EnableSyntaxChecking() *Root {
	return &Root{
		statements:     g.statements,
		gofmt:          g.gofmt,
		gofmtOptions:   g.gofmtOptions,
		goimports:      g.goimports,
		syntaxChecking: true,
	}
}

// Generate generates golang code according to registered statements.
func (g *Root) Generate(indentLevel int) (string, error) {
	generatedCode := ""

	for _, statement := range g.statements {
		gen, err := statement.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		generatedCode += gen
	}

	if g.syntaxChecking {
		_, err := g.applyGofmt(generatedCode, "-e")
		if err != nil {
			return "", err
		}
	}

	if g.gofmt {
		var err error
		generatedCode, err = g.applyGofmt(generatedCode, g.gofmtOptions...)
		if err != nil {
			return "", err
		}
	}

	if g.goimports {
		var err error
		generatedCode, err = g.applyGoimports(generatedCode)
		if err != nil {
			return "", err
		}
	}

	return generatedCode, nil
}

func (g *Root) applyGofmt(generatedCode string, gofmtOptions ...string) (string, error) {
	return applyCodeFormatter(generatedCode, "gofmt", gofmtOptions...)
}

func (g *Root) applyGoimports(generatedCode string) (string, error) {
	return applyCodeFormatter(generatedCode, "goimports")
}

func applyCodeFormatter(generatedCode string, formatterCmdName string, formatterOpts ...string) (string, error) {
	echoCmd := exec.Command("echo", generatedCode)
	formatterCmd := exec.Command(formatterCmdName, formatterOpts...)

	r, w := io.Pipe()
	echoCmd.Stdout = w
	formatterCmd.Stdin = r

	var out, errout bytes.Buffer
	formatterCmd.Stdout = &out
	formatterCmd.Stderr = &errout

	echoCmd.Start()
	if err := formatterCmd.Start(); err != nil {
		cmds := []string{formatterCmdName}
		return "", errmsg.CodeFormatterError(strings.Join(append(cmds, formatterOpts...), " "), errout.String(), err)
	}
	echoCmd.Wait()
	w.Close()
	err := formatterCmd.Wait()
	if err != nil {
		cmds := []string{formatterCmdName}
		return "", errmsg.CodeFormatterError(strings.Join(append(cmds, formatterOpts...), " "), errout.String(), err)
	}

	return out.String(), err
}
