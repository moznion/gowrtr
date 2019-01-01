package gowrtr

type Package struct {
	Name string
}

func NewPackage(packageName string) *Package {
	return &Package{
		Name: packageName,
	}
}

func (p *Package) String() string {
	return "package " + p.Name
}
