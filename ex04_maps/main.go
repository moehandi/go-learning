package main

import "fmt"

// format map[key]valueType
// The key type can be any type for which the operations == and != are defined,
// like: (string, int, float)
// So arrays, slices and structs cannot be used as key type, but pointers and interface types can.
// The value type can be any type; by using the empty interface as type
func main() {
	var ourAges = make(map[string]int)
	ourAges["adry"] = 26
	ourAges["andi"] = 26
	ourAges["fauzi"] = 26
	ourAges["firman"] = 22
	ourAges["geby"] = 26
	fmt.Println("ourAges:", ourAges)
	fmt.Println("andi age is", ourAges["andi"]) // get andi's age return 26
	fmt.Println("ade age is", ourAges["ade"])   // get ade's age which no key and value, return 0

	// equivalent to
	ages := map[string]int{
		"adry":   26,
		"andi":   26,
		"fauzi":  26,
		"firman": 22,
		"gebby":  26,
	}
	fmt.Println("ages:", ages)

	// lets loop over ages
	for name, age := range ages {
		fmt.Printf("%s is %d \n", name, age)
	}

	age, ok := ages["andi"] // return 26 true
	fmt.Println("age, ok", age, ok)

	age2, ok2 := ages["boy"] // there is no boy
	if !ok2 {
		fmt.Println(age2, "not found")
	}

	// we will see often combined above
	if age3, ok3 := ages["boy"]; !ok3 {
		fmt.Println("there is no boy key", age3)
	}

	// in example above we must used age3, to ignore age3 print just ignore with _
	if _, ok3 := ages["boy"]; !ok3 {
		fmt.Println("there is no boy key")
	}

	// nil map, see there is nothing assignment '=', make
	var nilMap map[string]int
	fmt.Println("nilmap == nil ", nilMap == nil)    // "true"
	fmt.Println("len(nilMap) == 0", len(nilMap) == 0) // "true"
	//nilMap["ambo"] = 21        // panic: assignment to entry in nil map

	var notNilMap map[string]int
	notNilMap = make(map[string]int)
	notNilMap["ambo"] = 21        // not panic: assignment to entry in not nil map
	fmt.Println("notNilMap:", notNilMap)
}
