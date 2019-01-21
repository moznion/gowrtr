package generator

import (
	"github.com/moznion/gowrtr/internal/errmsg"
)

// AnonymousFunc represents a code generator for anonymous func.
type AnonymousFunc struct {
	goFunc                 bool
	anonymousFuncSignature *AnonymousFuncSignature
	statements             []Statement
	funcInvocation         *FuncInvocation
	caller                 string
}

// NewAnonymousFunc returns a new `AnonymousFunc`.
// If `goFunc` is true, the anonymous function will be `go func`.
func NewAnonymousFunc(goFunc bool, signature *AnonymousFuncSignature, statements ...Statement) *AnonymousFunc {
	return &AnonymousFunc{
		goFunc:                 goFunc,
		anonymousFuncSignature: signature,
		statements:             statements,
		caller:                 fetchClientCallerLine(),
	}
}

// AddStatements adds statements for the function to `AnonymousFunc`. This does *not* set, just add.
// This method returns a *new* `AnonymousFunc`; it means this method acts as immutable.
func (ifg *AnonymousFunc) AddStatements(statements ...Statement) *AnonymousFunc {
	return &AnonymousFunc{
		goFunc:                 ifg.goFunc,
		anonymousFuncSignature: ifg.anonymousFuncSignature,
		statements:             append(ifg.statements, statements...),
		funcInvocation:         ifg.funcInvocation,
		caller:                 ifg.caller,
	}
}

// Statements sets statements for the function to `AnonymousFunc`. This does *not* add, just set.
// This method returns a *new* `AnonymousFunc`; it means this method acts as immutable.
func (ifg *AnonymousFunc) Statements(statements ...Statement) *AnonymousFunc {
	return &AnonymousFunc{
		goFunc:                 ifg.goFunc,
		anonymousFuncSignature: ifg.anonymousFuncSignature,
		statements:             statements,
		funcInvocation:         ifg.funcInvocation,
		caller:                 ifg.caller,
	}
}

// Invocation sets an invocation of the anonymous func to `AnonymousFunc`.
// This method returns a *new* `AnonymousFunc`; it means this method acts as immutable.
func (ifg *AnonymousFunc) Invocation(funcInvocation *FuncInvocation) *AnonymousFunc {
	return &AnonymousFunc{
		goFunc:                 ifg.goFunc,
		anonymousFuncSignature: ifg.anonymousFuncSignature,
		statements:             ifg.statements,
		funcInvocation:         funcInvocation,
		caller:                 ifg.caller,
	}
}

// Generate generates an anonymous func as golang code.
func (ifg *AnonymousFunc) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := indent
	if ifg.goFunc {
		stmt += "go "
	}
	stmt += "func"

	if ifg.anonymousFuncSignature == nil {
		return "", errmsg.AnonymousFuncSignatureIsNilError(ifg.caller)
	}

	sig, err := ifg.anonymousFuncSignature.Generate(0)
	if err != nil {
		return "", err
	}
	stmt += sig + " {\n"

	nextIndentLevel := indentLevel + 1
	for _, generator := range ifg.statements {
		gen, err := generator.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += indent + "}"

	if funcInvocation := ifg.funcInvocation; funcInvocation != nil {
		invocation, err := funcInvocation.Generate(0)
		if err != nil {
			return "", err
		}
		stmt += invocation
	}

	stmt += "\n"

	return stmt, nil
}
