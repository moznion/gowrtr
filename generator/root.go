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

// Gofmt enables `gofmt`. If `gofmt` is enabled, it applies `gofmt` on code generation phase.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) Gofmt(gofmtOptions ...string) *Root {
	return &Root{
		statements:     g.statements,
		gofmt:          true,
		gofmtOptions:   gofmtOptions,
		goimports:      g.goimports,
		syntaxChecking: g.syntaxChecking,
	}
}

// Goimports enables `goimports`. If `goimports` is enabled, it applies `goimports` on code generation phase.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) Goimports() *Root {
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
	formatterCmd := exec.Command(formatterCmdName, formatterOpts...)
	stdinPipe, _ := formatterCmd.StdinPipe()

	var out, errout bytes.Buffer
	formatterCmd.Stdout = &out
	formatterCmd.Stderr = &errout

	err := formatterCmd.Start()
	if err != nil {
		cmds := []string{formatterCmdName}
		return "", errmsg.CodeFormatterError(strings.Join(append(cmds, formatterOpts...), " "), errout.String(), err)
	}

	_, err = io.WriteString(stdinPipe, generatedCode)
	if err != nil {
		return "", err
	}
	err = stdinPipe.Close()
	if err != nil {
		return "", err
	}

	err = formatterCmd.Wait()
	if err != nil {
		cmds := []string{formatterCmdName}
		return "", errmsg.CodeFormatterError(strings.Join(append(cmds, formatterOpts...), " "), errout.String(), err)
	}

	return out.String(), err
}
