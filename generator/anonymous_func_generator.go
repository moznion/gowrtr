package generator

import "github.com/moznion/gowrtr/internal/errmsg"

// AnonymousFuncGenerator represents a code generator for anonymous func.
type AnonymousFuncGenerator struct {
	GoFunc                 bool
	AnonymousFuncSignature *AnonymousFuncSignatureGenerator
	Statements             []StatementGenerator
	FuncInvocation         *FuncInvocationGenerator
}

// NewAnonymousFuncGenerator returns a new `AnonymousFuncGenerator`.
// If `goFunc` is true, the anonymous function will be `go func`.
func NewAnonymousFuncGenerator(goFunc bool, signature *AnonymousFuncSignatureGenerator, statements ...StatementGenerator) *AnonymousFuncGenerator {
	return &AnonymousFuncGenerator{
		GoFunc:                 goFunc,
		AnonymousFuncSignature: signature,
		Statements:             statements,
	}
}

// AddStatements adds statements for the function to `AnonymousFuncGenerator`.
// This method returns a *new* `AnonymousFuncGenerator`; it means this method acts as immutable.
func (ifg *AnonymousFuncGenerator) AddStatements(statements ...StatementGenerator) *AnonymousFuncGenerator {
	return &AnonymousFuncGenerator{
		GoFunc:                 ifg.GoFunc,
		AnonymousFuncSignature: ifg.AnonymousFuncSignature,
		Statements:             append(ifg.Statements, statements...),
		FuncInvocation:         ifg.FuncInvocation,
	}
}

// SetFuncInvocation sets an invocation of the anonymous func to `AnonymousFuncGenerator`.
// This method returns a *new* `AnonymousFuncGenerator`; it means this method acts as immutable.
func (ifg *AnonymousFuncGenerator) SetFuncInvocation(funcInvocation *FuncInvocationGenerator) *AnonymousFuncGenerator {
	return &AnonymousFuncGenerator{
		GoFunc:                 ifg.GoFunc,
		AnonymousFuncSignature: ifg.AnonymousFuncSignature,
		Statements:             ifg.Statements,
		FuncInvocation:         funcInvocation,
	}
}

// Generate generates an anonymous func as golang code.
func (ifg *AnonymousFuncGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := indent
	if ifg.GoFunc {
		stmt += "go "
	}
	stmt += "func"

	if ifg.AnonymousFuncSignature == nil {
		return "", errmsg.AnonymousFuncSignatureIsNilError()
	}

	sig, err := ifg.AnonymousFuncSignature.Generate(0)
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
