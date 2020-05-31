package generator

import (
	"github.com/moznion/gowrtr/internal/errmsg"
)

// Func represents a code generator for the func.
type Func struct {
	funcReceiver  *FuncReceiver
	funcSignature *FuncSignature
	statements    []Statement
	caller        string
}

// NewFunc returns a new `Func`.
func NewFunc(receiver *FuncReceiver, signature *FuncSignature, statements ...Statement) *Func {
	return &Func{
		funcReceiver:  receiver,
		funcSignature: signature,
		statements:    statements,
		caller:        fetchClientCallerLine(),
	}
}

// AddStatements adds statements for the func to `Func`. This does *not* set, just add.
// This method returns a *new* `Func`; it means this method acts as immutable.
func (fg *Func) AddStatements(statements ...Statement) *Func {
	return &Func{
		funcReceiver:  fg.funcReceiver,
		funcSignature: fg.funcSignature,
		statements:    append(fg.statements, statements...),
	}
}

// Statements sets statements for the func to `Func`. This does *not* add, just set.
// This method returns a *new* `Func`; it means this method acts as immutable.
func (fg *Func) Statements(statements ...Statement) *Func {
	return &Func{
		funcReceiver:  fg.funcReceiver,
		funcSignature: fg.funcSignature,
		statements:    statements,
	}
}

// Generate generates a func block as golang code.
func (fg *Func) Generate(indentLevel int) (string, error) {
	indent := BuildIndent(indentLevel)

	stmt := indent + "func "

	receiver := ""
	if fg.funcReceiver != nil {
		var err error
		receiver, err = fg.funcReceiver.Generate(0)
		if err != nil {
			return "", err
		}
	}
	if receiver != "" {
		stmt += receiver + " "
	}

	if fg.funcSignature == nil {
		return "", errmsg.FuncSignatureIsNilError(fg.caller)
	}
	sig, err := fg.funcSignature.Generate(0)
	if err != nil {
		return "", err
	}
	stmt += sig + " {\n"

	nextIndentLevel := indentLevel + 1
	for _, c := range fg.statements {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += indent + "}\n"

	return stmt, nil
}
