package generator

import (
	"fmt"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// If represents a code generator for `if`, `else-if` and `else` block.
type If struct {
	condition    string
	statements   []Statement
	elseIfBlocks []*ElseIf
	elseBlock    *Else
	caller       string
}

// NewIf returns a new `If`.
func NewIf(condition string, statements ...Statement) *If {
	return &If{
		condition:  condition,
		statements: statements,
		caller:     fetchClientCallerLine(),
	}
}

// AddStatements adds statements for `if` block to `If`. This does *not* set, just add.
// This method returns a *new* `If`; it means this method acts as immutable.
func (ig *If) AddStatements(statements ...Statement) *If {
	return &If{
		condition:    ig.condition,
		statements:   append(ig.statements, statements...),
		elseIfBlocks: ig.elseIfBlocks,
		elseBlock:    ig.elseBlock,
		caller:       ig.caller,
	}
}

// Statements sets statements for `if` block to `If`. This does *not* add, just set.
// This method returns a *new* `If`; it means this method acts as immutable.
func (ig *If) Statements(statements ...Statement) *If {
	return &If{
		condition:    ig.condition,
		statements:   statements,
		elseIfBlocks: ig.elseIfBlocks,
		elseBlock:    ig.elseBlock,
		caller:       ig.caller,
	}
}

// AddElseIf adds `else-if` block to `If`. This does *not* set, just add.
// This method returns a *new* `If`; it means this method acts as immutable.
func (ig *If) AddElseIf(blocks ...*ElseIf) *If {
	return &If{
		condition:    ig.condition,
		statements:   ig.statements,
		elseIfBlocks: append(ig.elseIfBlocks, blocks...),
		elseBlock:    ig.elseBlock,
		caller:       ig.caller,
	}
}

// ElseIf sets `else-if` block to `If`. This does *not* add, just set.
// This method returns a *new* `If`; it means this method acts as immutable.
func (ig *If) ElseIf(blocks ...*ElseIf) *If {
	return &If{
		condition:    ig.condition,
		statements:   ig.statements,
		elseIfBlocks: blocks,
		elseBlock:    ig.elseBlock,
		caller:       ig.caller,
	}
}

// Else sets `else` block to `If`.
// This method returns a *new* `If`; it means this method acts as immutable.
func (ig *If) Else(block *Else) *If {
	return &If{
		condition:    ig.condition,
		statements:   ig.statements,
		elseIfBlocks: ig.elseIfBlocks,
		elseBlock:    block,
		caller:       ig.caller,
	}
}

// Generate generates `if` block as golang code.
func (ig *If) Generate(indentLevel int) (string, error) {
	indent := BuildIndent(indentLevel)

	if ig.condition == "" {
		return "", errmsg.IfConditionIsEmptyError(ig.caller)
	}

	stmt := fmt.Sprintf("%sif %s {\n", indent, ig.condition)

	nextIndentLevel := indentLevel + 1
	for _, c := range ig.statements {
		gen, err := c.Generate(nextIndentLevel)
		if err != nil {
			return "", err
		}
		stmt += gen
	}

	stmt += fmt.Sprintf("%s}", indent)

	for _, elseIfBlock := range ig.elseIfBlocks {
		if elseIfBlock == nil {
			continue
		}

		elseIfCode, err := elseIfBlock.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		stmt += elseIfCode
	}

	if elseBlock := ig.elseBlock; elseBlock != nil {
		elseCode, err := elseBlock.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		stmt += elseCode
	}

	stmt += "\n"

	return stmt, nil
}
