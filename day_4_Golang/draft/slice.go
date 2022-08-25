package main

import "fmt"

func main() {
	// var {nama-variabel} [{kapasitas}]{type-data}
	// [...] variadic alias elastis arraynya, bebas diisi brppun
	var buah4 = [4]string{"apple", "nanas", "melon", "alpukat"}
	var buahUnlimited = [...]string{"apple", "nanas", "melon", "alpukat", "durian", "mangga"}

	fmt.Println("buah: ", buah4)
	fmt.Println("buah: ", buahUnlimited)
}
