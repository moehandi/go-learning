package main

func main()  {
	returnX, returnY := getAddOne(7, 7)
	println("unnamed return", returnX, returnY)

	returnX2, returnY2 := getAddTwo(8, 8)
	println("named return", returnX2, returnY2)
}

// multiple return value both named or unnamed if return more than 1 must be enclosed by ()
// Unnamed return values
func getAddOne(x, y int) (int, int)  {
	return x + 1, y + 1
}

// Named return values (xOut and yOut)
func getAddTwo(x, y int) (xOut, yOut int) {
	xOut = x + 2
	yOut = y + 2
	return
}
