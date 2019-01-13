package generator

import (
	"fmt"
)

// IfGenerator represents a code generator for `if`, `else-if` and `else` block.
type IfGenerator struct {
	Condition    string
	Statements   []StatementGenerator
	ElseIfBlocks []*ElseIfGenerator
	ElseBlock    *ElseGenerator
}

// NewIfGenerator returns a new `IfGenerator`.
func NewIfGenerator(condition string, statements ...StatementGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:  condition,
		Statements: statements,
	}
}

// AddStatements adds statements for `if` block to `IfGenerator`.
// This method returns a *new* `IfGenerator`; it means this method acts as immutable.
func (ig *IfGenerator) AddStatements(statements ...StatementGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:    ig.Condition,
		Statements:   append(ig.Statements, statements...),
		ElseIfBlocks: ig.ElseIfBlocks,
		ElseBlock:    ig.ElseBlock,
	}
}

// AddElseIfBlocks adds `else-if` block to `IfGenerator`.
// This method returns a *new* `IfGenerator`; it means this method acts as immutable.
func (ig *IfGenerator) AddElseIfBlocks(blocks ...*ElseIfGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:    ig.Condition,
		Statements:   ig.Statements,
		ElseIfBlocks: append(ig.ElseIfBlocks, blocks...),
		ElseBlock:    ig.ElseBlock,
	}
}

// SetElseBlock sets `else` block to `IfGenerator`.
// This method returns a *new* `IfGenerator`; it means this method acts as immutable.
func (ig *IfGenerator) SetElseBlock(block *ElseGenerator) *IfGenerator {
	return &IfGenerator{
		Condition:    ig.Condition,
		Statements:   ig.Statements,
		ElseIfBlocks: ig.ElseIfBlocks,
		ElseBlock:    block,
	}
}

// Generate generates `if` block as golang code.
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
