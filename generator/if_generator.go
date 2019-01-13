package generator

import (
	"fmt"
)

type IfGenerator struct {
	Condition    string
	Statements   []StatementGenerator
	ElseIfBlocks []*ElseIfGenerator
	ElseBlock    *ElseGenerator
}

func NewIfGenerator(condition string, statements ...StatementGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:  condition,
		Statements: statements,
	}
}

func (ig *IfGenerator) AddStatements(statements ...StatementGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:    ig.Condition,
		Statements:   append(ig.Statements, statements...),
		ElseIfBlocks: ig.ElseIfBlocks,
		ElseBlock:    ig.ElseBlock,
	}
}

func (ig *IfGenerator) AddElseIfBlocks(blocks ...*ElseIfGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:    ig.Condition,
		Statements:   ig.Statements,
		ElseIfBlocks: append(ig.ElseIfBlocks, blocks...),
		ElseBlock:    ig.ElseBlock,
	}
}

func (ig *IfGenerator) SetElseBlock(block *ElseGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:    ig.Condition,
		Statements:   ig.Statements,
		ElseIfBlocks: ig.ElseIfBlocks,
		ElseBlock:    block,
	}
}

func (ig *IfGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	stmt := fmt.Sprintf("%sif %s {\n", indent, ig.Condition)

	nextIndentLevel := indentLevel + 1
	for _, c := range ig.Statements {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += fmt.Sprintf("%s}", indent)

	for _, elseIfBlock := range ig.ElseIfBlocks {
		if elseIfBlock == nil {
			continue
		}

		elseIfCode, err := elseIfBlock.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		stmt += elseIfCode
	}

	if elseBlock := ig.ElseBlock; elseBlock != nil {
		elseCode, err := elseBlock.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		stmt += elseCode
	}

	stmt += "\n"

	return stmt, nil
}
