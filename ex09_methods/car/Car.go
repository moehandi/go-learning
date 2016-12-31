package car

type Car struct {
	color string // lower case, it's unexported or private in OO
	door int
}

// it's like getter, in Go use directyly Color instead GetColor
func (c *Car) Color() string {
	return c.color
}

// it's like setter
func (c *Car) SetColor(newColor string)  {
	c.color = newColor
}

func (c *Car) Door() int {
	return c.door
}

// it's like setter
func (c *Car) SetDoor(newDoor int)  {
	c.door = newDoor
}