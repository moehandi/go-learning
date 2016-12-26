package main

import "fmt"

func main() {
	start()
	defer close() //executes after all function done
	ongoing()
}

func start()  {
	fmt.Println("start")
}

func close() {
	fmt.Println("close")
}

func ongoing() {
	fmt.Println("on going")
}
