package main

import "fmt"

func main() {
	// var {nama-variabel} [{kapasitas}]{type-data}
	// [...] variadic alias elastis arraynya, bebas diisi brppun
	var buah = []string{"apple", "nanas", "melon", "alpukat"}

	buahSebagian := buah[2:4]

	fmt.Println("buah: ", buah)
	fmt.Println("buah Sebagian: ", buahSebagian)
}
