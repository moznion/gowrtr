package generator

// Newline represents a code generator for newline character.
type Newline struct {
}

// NewNewline returns a new `Newline`.
func NewNewline() *Newline {
	return &Newline{}
}

// Generate generates a newline statement as golang code.
func (n *Newline) Generate(indentLevel int) (string, error) {
	return "\n", nil
}
