package main

import "fmt"

type Katalog struct {
	id    int
	name  string
	price int
}

// Mencari mahal dan murah

func murahMahal(katalog []Katalog) (min Katalog, max Katalog) {
	min = katalog[0]
	max = katalog[0]
	for _, value := range katalog {
		if value.price < min.price {
			min = value
		}
		if value.price > max.price {
			max = value
		}
	}
	return
}

// mencari barang dengan harga 10000

func isSepuluh(katalog Katalog) bool {
	return katalog.price == 10000
}

func main() {
	// var s1 katalog
	// s1.name = "Zhang Purnama"

	katalogs := []Katalog{
		{id: 1, name: "Benih Lele", price: 50000},
		{id: 2, name: "Pakan lele cap menara", price: 25000},
		{id: 3, name: "Probiotik A", price: 175000},
		{id: 4, name: "Probiotik nila B", price: 10000},
		{id: 5, name: "Pakan nila", price: 20000},
		{id: 6, name: "Benih nila", price: 20000},
		{id: 7, name: "Cupang", price: 5000},
		{id: 8, name: "Benih Nila", price: 30000},
		{id: 9, name: "Benih Cupang", price: 10000},
		{id: 10, name: "Probiotik B", price: 10000},
	}

	// var murahmahal []Katalog
	// Mencari mahal dan murah
	murah, mahal := murahMahal(katalogs)
	fmt.Println("Barang termurah: ", murah)
	fmt.Println("Barang termahal: ", mahal)

	// mencari barang dengan harga 10000
	var lists []Katalog

	for _, katalog := range katalogs {
		if isSepuluh(katalog) {
			lists = append(lists, katalog)
		}

	}

	fmt.Println("Harga 10000:")
	for _, price := range lists {
		fmt.Println(price)
	}
	// for _, katalog := range katalogs {
	// 	fmt.Println("-------------")
	// 	fmt.Println("ID: ", katalog.id)
	// 	fmt.Println("Barang: ", katalog.name)
	// 	fmt.Println("Harga: ", katalog.price)

	// }

}
