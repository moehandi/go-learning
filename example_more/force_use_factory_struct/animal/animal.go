package animal

type animal struct {
	Name string
	Species string
}

func NewAnimal(name, species string) *animal  {
	a := &animal{name, species}
	return a
}
