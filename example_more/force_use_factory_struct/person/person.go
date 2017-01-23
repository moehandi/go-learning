package person

// lower case person cannot access directly from outside
type person struct {
	Name string
	Age int
}

// it's like constructor
func New(name string, age int) *person {
	p := new(person)
	p.Name = name
	p.Age = age
	return p
}