package gowrtr

import (
	"github.com/moznion/gowrtr/errmsg"
)

type FuncGenerator struct {
	FuncReceiver  *FuncReceiverGenerator
	FuncSignature *FuncSignatureGenerator
	Generators    []CodeGenerator
}

func NewFuncGenerator(receiver *FuncReceiverGenerator, signature *FuncSignatureGenerator, generators ...CodeGenerator) *FuncGenerator {
	return &FuncGenerator{
		FuncReceiver:  receiver,
		FuncSignature: signature,
		Generators:    generators,
	}
}

func (fg *FuncGenerator) AddStatements(generators ...CodeGenerator) *FuncGenerator {
	return &FuncGenerator{
		FuncReceiver:  fg.FuncReceiver,
		FuncSignature: fg.FuncSignature,
		Generators:    append(fg.Generators, generators...),
	}
}

func (fg *FuncGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := indent + "func "

	receiver := ""
	if fg.FuncReceiver != nil {
		var err error
		receiver, err = fg.FuncReceiver.Generate(0)
		if err != nil {
			return "", err
		}
	}
	if receiver != "" {
		stmt += receiver + " "
	}

	if fg.FuncSignature == nil {
		return "", errmsg.FuncSignatureIsNilError()
	}
	sig, err := fg.FuncSignature.Generate(0)
	if err != nil {
		return "", err
	}
	stmt += sig + " {\n"

	nextIndentLevel := indentLevel + 1
	for _, c := range fg.Generators {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += indent + "}\n"

	return stmt, nil
}
