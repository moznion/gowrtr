package gowrtr

import "github.com/moznion/gowrtr/errmsg"

type InlineFuncGenerator struct {
	GoFunc              bool
	InlineFuncSignature *InlineFuncSignatureGenerator
	Generators          []CodeGenerator
	FuncInvocation      *FuncInvocationGenerator
}

func NewInlineFuncGenerator(goFunc bool, signature *InlineFuncSignatureGenerator, generators ...CodeGenerator) *InlineFuncGenerator {
	return &InlineFuncGenerator{
		GoFunc:              goFunc,
		InlineFuncSignature: signature,
		Generators:          generators,
	}
}

func (ifg *InlineFuncGenerator) AddStatements(generators ...CodeGenerator) *InlineFuncGenerator {
	return &InlineFuncGenerator{
		GoFunc:              ifg.GoFunc,
		InlineFuncSignature: ifg.InlineFuncSignature,
		Generators:          append(ifg.Generators, generators...),
		FuncInvocation:      ifg.FuncInvocation,
	}
}

func (ifg *InlineFuncGenerator) AddFuncInvocation(funcInvocation *FuncInvocationGenerator) *InlineFuncGenerator {
	return &InlineFuncGenerator{
		GoFunc:              ifg.GoFunc,
		InlineFuncSignature: ifg.InlineFuncSignature,
		Generators:          ifg.Generators,
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
	for _, generator := range ifg.Generators {
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
