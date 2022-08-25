package main

import "fmt"

// define constant sendiri diluar main
const (
	errorMessage = "error"
	infoMessage  = "info"
)

// func {nama-function}({parameter}){return-type-data}
// func swap(a string, b string){

func hello(nama string) string {
	return "hallo, " + nama
}

// func hello(nama string) string {
// 	return "hallo, " + nama + ". Anda berumur " + int
// }

// func swap
func swap(a string, b string) (string, string) {
	return a, b
}

// func helloVariadic
func helloVariadic(name ...string) string {
	msg := "hello, "
	for _, v := range name {
		msg += v
	}
	return msg
}

// struc
/// struc student
type student struct {
	name    string
	grade   int
	sekolah school //embeded struct
}

/// struc school
type school struct {
	name    string
	address string
	kelas   class
}

type class struct {
	name string
}

func main() {

	// var {nama variabel} {type data}
	var decimalNumber float64 = 3.1

	// var {nama variabel}
	var decimalNumber2 = 3.2

	// shorthand
	decimalNumber3 := 3.3

	// define multi variable

	// var {nama variabel} ... {type data}
	var decimalNumber4, decimalNumber5 float64

	decimalNumber4 = 3.4
	decimalNumber5 = 3.5

	varA, varB := "nilai A", "nilai B"
	fmt.Println("var A before: ", varA)
	fmt.Println("var B before: ", varB)

	varB, varA = varA, varB
	fmt.Println("var A after: ", varA)
	fmt.Println("var B after: ", varB)

	fmt.Printf("decimal number menggunakan formating: %f\n", decimalNumber)
	fmt.Println("decimal number: ", decimalNumber)
	fmt.Println("decimal number 2: ", decimalNumber2)
	fmt.Println("decimal number 3: ", decimalNumber3)
	fmt.Println("decimal number 4: ", decimalNumber4)
	fmt.Println("decimal number 5: ", decimalNumber5)
	fmt.Println("var A: ", varA)
	fmt.Println("var B: ", varB)
	fmt.Println("hello world")

	// Boolean
	var isExist bool = true
	fmt.Println("isExist: ", isExist)

	// string
	// const firstName string = "catto" //const gabisa diganti variabelnya
	var firstName string = "catto"
	var lastName = "meow"

	var name = firstName + lastName

	fmt.Println("Nama: ", name)

	fmt.Println("value error: ", errorMessage)
	fmt.Println("value info: ", infoMessage)

	//else if blablabla
	var point = 8

	if point == 10 {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}

	// for range
	var buah = [2]string{"apel", "nanas"}

	for i, value := range buah {
		fmt.Println("index: ", i)
		fmt.Println("index: ", value)
		fmt.Println("---------------")
	}

	// for loop break
	for i := 0; i < 10; i++ {
		if i == 2 { //skip angka 2, karena i = 2
			continue //artinya skip code dibawahnya
		}
		if i == 8 {
			break //berhenti total
		}
		fmt.Println("number: ", i)
	}

	//for loop biasa
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("number: ", i)
	// }

	// penggunaan fungsi hello()
	hello := hello("cattoMeow")
	fmt.Println(hello)

	// penggunaan fungsi swap
	// a, b := swap("varA", "varB")
	// fmt.Println("var a:", a)
	// fmt.Println("var b:", b)

	///swap hanya pakai variabel a saja
	a, _ := swap("varA", "varB")
	fmt.Println("var a:", a)

	// func helloVariadic()
	hello1 := helloVariadic("cwtto ", "meow ", "github") //tipe data harus sama, else error
	fmt.Println("hello1: ", hello1)

	// struct
	var s1 student
	s1.name = "Zhang Purnama"
	s1.grade = 2
	// embeded struct
	s1.sekolah.name = "Tokyo University"
	s1.sekolah.address = "Japan"

	students := []student{
		{name: "Zhang", grade: 10, sekolah: school{name: "Tokyo University", address: "Japan", kelas: class{name: "Kelas 10"}}},
		{name: "Purnama", grade: 8, sekolah: school{name: "Tokyo University", address: "Japan", kelas: class{name: "Kelas 12"}}},
	}

	for _, student := range students {
		fmt.Println("-------------")
		fmt.Println("name: ", student.name)
		fmt.Println("grade: ", student.grade)
		fmt.Println("sekolah: ", student.sekolah.name)
		fmt.Println("alamat: ", student.sekolah.address)
		fmt.Println("kelas: ", student.sekolah.kelas.name)
		fmt.Println("-------------")
	}

	fmt.Println("NAME: ", s1.name)
	fmt.Println("GRADE: ", s1.grade)
	fmt.Println("SEKOLAH: ", s1.sekolah.name)
	fmt.Println("ADDRESS: ", s1.sekolah.address)
}
