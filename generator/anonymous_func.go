package generator

import "github.com/moznion/gowrtr/internal/errmsg"

// AnonymousFunc represents a code generator for anonymous func.
type AnonymousFunc struct {
	GoFunc                 bool
	AnonymousFuncSignature *AnonymousFuncSignature
	Statements             []Statement
	FuncInvocation         *FuncInvocation
}

// NewAnonymousFunc returns a new `AnonymousFunc`.
// If `goFunc` is true, the anonymous function will be `go func`.
func NewAnonymousFunc(goFunc bool, signature *AnonymousFuncSignature, statements ...Statement) *AnonymousFunc {
	return &AnonymousFunc{
		GoFunc:                 goFunc,
		AnonymousFuncSignature: signature,
		Statements:             statements,
	}
}

// AddStatements adds statements for the function to `AnonymousFunc`.
// This method returns a *new* `AnonymousFunc`; it means this method acts as immutable.
func (ifg *AnonymousFunc) AddStatements(statements ...Statement) *AnonymousFunc {
	return &AnonymousFunc{
		GoFunc:                 ifg.GoFunc,
		AnonymousFuncSignature: ifg.AnonymousFuncSignature,
		Statements:             append(ifg.Statements, statements...),
		FuncInvocation:         ifg.FuncInvocation,
	}
}

// SetFuncInvocation sets an invocation of the anonymous func to `AnonymousFunc`.
// This method returns a *new* `AnonymousFunc`; it means this method acts as immutable.
func (ifg *AnonymousFunc) SetFuncInvocation(funcInvocation *FuncInvocation) *AnonymousFunc {
	return &AnonymousFunc{
		GoFunc:                 ifg.GoFunc,
		AnonymousFuncSignature: ifg.AnonymousFuncSignature,
		Statements:             ifg.Statements,
		FuncInvocation:         funcInvocation,
	}
}

// Generate generates an anonymous func as golang code.
func (ifg *AnonymousFunc) Generate(indentLevel int) (string, error) {
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
