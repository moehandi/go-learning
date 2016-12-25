package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Address string   `json:"address,omitempty"`
	Active  bool     `json:"active"`
	Skill   []string `json:"skill"`
}

func main() {
	userA := User {1, "moehandi", "Jakarta", true, []string{"java","golang","android","nodejs"}}
	userData := []User{
		{1, "adry", "Bandung", true, []string{"ruby","polymer","mongodb","redis"}},
		{2, "geby", "Riau", true, []string{"java","android"}},
		{3, "moehandi", "Bengkulu", true, []string{"java","golang","android","nodejs"}},
		{4, "raka", "Jakarta", true, []string{"java","android"}},
	}
	// Marshalling from Data to JSON
	// Marshal produces a byte slice containing a very long string with no extraneous white space
	user, err := json.Marshal(userA) // Marshalling userA instance or Object
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	fmt.Printf("%s\n", user)

	// Marshalling userData Array of User
	//data, err := json.Marshal(userData) // Marshalling userData array
	data, err := json.MarshalIndent(userData,""," ") // Marshalling userData array with indent
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	fmt.Printf("%s\n", data)

	// TODO Unmarshall : is From JSON to Data(structs)
}
