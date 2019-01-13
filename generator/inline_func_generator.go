package gowrtr

import "github.com/moznion/gowrtr/internal/errmsg"

type InlineFuncGenerator struct {
	GoFunc              bool
	InlineFuncSignature *InlineFuncSignatureGenerator
	Statements          []CodeGeneratable
	FuncInvocation      *FuncInvocationGenerator
}

func NewInlineFuncGenerator(goFunc bool, signature *InlineFuncSignatureGenerator, statements ...CodeGeneratable) *InlineFuncGenerator {
	return &InlineFuncGenerator{
		GoFunc:              goFunc,
		InlineFuncSignature: signature,
		Statements:          statements,
	}
}

func (ifg *InlineFuncGenerator) AddStatements(statements ...CodeGeneratable) *InlineFuncGenerator {
	return &InlineFuncGenerator{
		GoFunc:              ifg.GoFunc,
		InlineFuncSignature: ifg.InlineFuncSignature,
		Statements:          append(ifg.Statements, statements...),
		FuncInvocation:      ifg.FuncInvocation,
	}
}

func (ifg *InlineFuncGenerator) AddFuncInvocation(funcInvocation *FuncInvocationGenerator) *InlineFuncGenerator {
	return &InlineFuncGenerator{
		GoFunc:              ifg.GoFunc,
		InlineFuncSignature: ifg.InlineFuncSignature,
		Statements:          ifg.Statements,
		FuncInvocation:      funcInvocation,
	}
}

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
