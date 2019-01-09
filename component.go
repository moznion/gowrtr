package gowrtr

type Component interface {
	GenerateCode() (string, error)
}
