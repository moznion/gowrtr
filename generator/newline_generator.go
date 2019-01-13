package generator

// NewlineGenerator represents a code generator for newline character.
type NewlineGenerator struct {
}

// NewNewlineGenerator returns a new `NewlineGenerator`.
func NewNewlineGenerator() *NewlineGenerator {
	return &NewlineGenerator{}
}

// Generate generates a newline statement as golang code.
func (n *NewlineGenerator) Generate(indentLevel int) (string, error) {
	return "\n", nil
}
