package main

import "fmt"

type Mixer interface { Mix() string }

type Mixo struct { capacity int }

type Ximo struct { capacity int }

func main() {
	var mix Mixer

	mixo := new(Mixo)
	mixo.capacity = 5
	mix = mixo

	if m, ok := mix.(*Mixo); ok {
		fmt.Printf("The type of mix is: %T\n", m)
	}

	if x, ok := mix.(*Ximo); ok {
		fmt.Printf("The type of mix is: %T\n", x)
	} else {
		fmt.Println("a mix doesn't contain variable of type Ximo")
	}
}

func (m *Mixo) Mix() string {
	return string(m.capacity * 2)
}

func (x *Ximo) Mix() string {
	return string(x.capacity * 8)
}



