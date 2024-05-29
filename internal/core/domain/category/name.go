package category

type Name string

func NewName(name string) (Name, error) {
	return Name(name), nil
}
