package main

import "fmt"

func main()  {
	mapX := map[string]int{
		"adry":   26,
		"andi":   26,
		"fauzi":  26,
		"firman": 22,
		"gebby":  26,
	}

	mapX2 := map[string]int{
		"adry":   26,
		"andi":   26,
		"fauzi":  26,
		"firman": 22,
		"gebby":  26,
	}

	mapY := map[string]int{
		"adry":   26,
		"andi":   26,
		"firman": 22,
		"gebby":  26,
	}

	fmt.Println("mapX equals mapX2", equals(mapX, mapX2))
	fmt.Println("mapX equals mapY", equals(mapX, mapY)) // because fauzi isn't in mapY

}

func equals(x, y map[string] int) bool {
	if len(x) != len(y) {
		return false
	}

	for key, xValue := range x {
		if yValue, ok := y[key]; !ok || yValue != xValue {
			return false
		}
	}

	return true
}
