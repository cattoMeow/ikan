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

}
