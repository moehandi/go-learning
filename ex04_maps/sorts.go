package main

import (
	"fmt"
	"sort"
)

// Lets try to sorting unordered map with it's key
func main()  {
	// unordered key, it's not alphabetical order
	ages := map[string]int{
		"fauzi":   26,
		"gebby":   26,
		"andi":  26,
		"firman": 22,
		"adry":  26,
	}

	fmt.Println("ages:", ages)

	// declare new slice string
	var names []string
	for name := range ages  {
		// append each name in key ages, key is name of people
		names = append(names, name)
	}
	// now sort all elements in slice
	sort.Strings(names)
	// loop the sorted name to print each name with it's age
	for _, name := range names {
		fmt.Printf("%s\t %d\n", name, ages[name])
	}
}
