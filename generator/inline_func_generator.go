package generator

import "github.com/moznion/gowrtr/internal/errmsg"

// InlineFuncGenerator represents a code generator for inline func.
type InlineFuncGenerator struct {
	GoFunc              bool
	InlineFuncSignature *InlineFuncSignatureGenerator
	Statements          []StatementGenerator
	FuncInvocation      *FuncInvocationGenerator
}

// NewInlineFuncGenerator returns a new `InlineFuncGenerator`.
// If `goFunc` is true, the inline function will be `go func`.
func NewInlineFuncGenerator(goFunc bool, signature *InlineFuncSignatureGenerator, statements ...StatementGenerator) *InlineFuncGenerator {
	return &InlineFuncGenerator{
		GoFunc:              goFunc,
		InlineFuncSignature: signature,
		Statements:          statements,
	}
}

// AddStatements adds statements for the function to `InlineFuncGenerator`.
// This method returns a *new* `InlineFuncGenerator`; it means this method acts as immutable.
func (ifg *InlineFuncGenerator) AddStatements(statements ...StatementGenerator) *InlineFuncGenerator {
	return &InlineFuncGenerator{
		GoFunc:              ifg.GoFunc,
		InlineFuncSignature: ifg.InlineFuncSignature,
		Statements:          append(ifg.Statements, statements...),
		FuncInvocation:      ifg.FuncInvocation,
	}
}

// SetFuncInvocation sets an invocation of the inline func to `InlineFuncGenerator`.
// This method returns a *new* `InlineFuncGenerator`; it means this method acts as immutable.
func (ifg *InlineFuncGenerator) SetFuncInvocation(funcInvocation *FuncInvocationGenerator) *InlineFuncGenerator {
	return &InlineFuncGenerator{
		GoFunc:              ifg.GoFunc,
		InlineFuncSignature: ifg.InlineFuncSignature,
		Statements:          ifg.Statements,
		FuncInvocation:      funcInvocation,
	}
}

// Generate generates an inline func as golang code.
func (ifg *InlineFuncGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := indent
	if ifg.GoFunc {
		stmt += "go "
	}
	stmt += "func"

	if ifg.InlineFuncSignature == nil {
		return "", errmsg.InlineFuncSignatureIsNilError()
	}

	sig, err := ifg.InlineFuncSignature.Generate(0)
	if err != nil {
		return "", err
	}
	stmt += sig + " {\n"

	nextIndentLevel := indentLevel + 1
	for _, generator := range ifg.Statements {
		gen, err := generator.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += indent + "}"

	if funcInvocation := ifg.FuncInvocation; funcInvocation != nil {
		invocation, err := funcInvocation.Generate(0)
		if err != nil {
			return "", err
		}
		stmt += invocation
	}

	stmt += "\n"

	return stmt, nil
}
