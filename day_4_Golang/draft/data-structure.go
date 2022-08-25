package main

import "fmt"

func main() {
	// data structure map

	// var {nama-variable} map[{type-key}]{tyoe-value}
	var student map[string]int
	student = map[string]int{}
	//[string] bisa diganti keynya jadi [int], sesuaikan key bawahnya
	student["zhang"] = 90
	student["purnama"] = 100

	fmt.Println("student: ", student)
	fmt.Println("student specific: ", student["zhang"])

	// array
	var buah4 = [4]string{"apple", "nanas", "melon", "alpukat"}
	var buahUnlimited = [...]string{"apple", "nanas", "melon", "alpukat", "durian", "mangga"}

	fmt.Println("buah: ", buah4)
	fmt.Println("buah: ", buahUnlimited)
}
