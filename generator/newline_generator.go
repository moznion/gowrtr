package gowrtr

type NewlineGenerator struct {
}

func NewNewlineGenerator() *NewlineGenerator {
	return &NewlineGenerator{}
}

func (n *NewlineGenerator) Generate(indentLevel int) (string, error) {
	return "\n", nil
}
