package main

import "fmt"

// Golang primitive types: numeric(int,float,complex), string, and boolean
func main() {
	var num1 int8 = 8 // int : int8, int16, int32(rune), int64
	var float1 float32 = 1.23 // float: float32 and float64
	var str1 string = "string1"
	var bool1 bool = true

	fmt.Println(num1, float1, str1, bool1)

	var num2 int8 = 8
	var num3 int16 = 8
	//num3 = num2 // it will fail because we can't assign num2 as int8 to num3 as int16
	num3 = int16(num2) // cast int8 to int16
	fmt.Println(num3)

	// byte type is uint8 type
	var intA uint8 = 28
	var byteA byte = 88
	intA = byteA // success because byte is uint8
	fmt.Println(intA)

	// rune type is int32
	var runeA rune = 1268
	var int32A int32 = 1268
	runeA = int32A
	fmt.Println(runeA)


}
