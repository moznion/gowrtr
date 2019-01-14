package generator

import (
	"fmt"
)

// If represents a code generator for `if`, `else-if` and `else` block.
type If struct {
	Condition    string
	Statements   []Statement
	ElseIfBlocks []*ElseIf
	ElseBlock    *Else
}

// NewIf returns a new `If`.
func NewIf(condition string, statements ...Statement) *If {
	return &If{
		Condition:  condition,
		Statements: statements,
	}
}

// AddStatements adds statements for `if` block to `If`.
// This method returns a *new* `If`; it means this method acts as immutable.
func (ig *If) AddStatements(statements ...Statement) *If {
	return &If{
		Condition:    ig.Condition,
		Statements:   append(ig.Statements, statements...),
		ElseIfBlocks: ig.ElseIfBlocks,
		ElseBlock:    ig.ElseBlock,
	}
}

// AddElseIfBlocks adds `else-if` block to `If`.
// This method returns a *new* `If`; it means this method acts as immutable.
func (ig *If) AddElseIfBlocks(blocks ...*ElseIf) *If {
	return &If{
		Condition:    ig.Condition,
		Statements:   ig.Statements,
		ElseIfBlocks: append(ig.ElseIfBlocks, blocks...),
		ElseBlock:    ig.ElseBlock,
	}
}

// SetElseBlock sets `else` block to `If`.
// This method returns a *new* `If`; it means this method acts as immutable.
func (ig *If) SetElseBlock(block *Else) *If {
	return &If{
		Condition:    ig.Condition,
		Statements:   ig.Statements,
		ElseIfBlocks: ig.ElseIfBlocks,
		ElseBlock:    block,
	}
}

// Generate generates `if` block as golang code.
func (ig *If) Generate(indentLevel int) (string, error) {
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
