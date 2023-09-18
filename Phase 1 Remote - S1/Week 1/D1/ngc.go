package main

import (
	"fmt"
)

func main() {
	// Variable 1

	// Soal A
	var myNum int32 = 50

	fmt.Println("My Number is", myNum)

	// Soal B
	var myNum2 float32 = 51
	fmt.Println("My Number is", myNum2)

	// Soal C
	var myNumStr string = "50"
	fmt.Println("My Number is", myNumStr)

	// Variable 2
	var x int32 = 5
	var y int32 = 10
	z := x + y
	fmt.Println("The sum of x and y is", z)

	// CLI

	// Membuat program CLI yang menerima input nama (string) dan akan menampilkan `hello [nama]`
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello", name)

	// Array & Slice 1
	people := []string{"Walt", "Jesse", "Skyler", "Saul"}
	fmt.Println("Length of people:", len(people))
	people = append(people, "Hank", "Marie")
	fmt.Println("Length of people:", len(people))
	fmt.Println(people)

	// Array & Slice 2
	// Membuat array dengan kapasitas 3 elemen
	data := []map[string]string{
		{"name": "Hank", "gender": "M"},
		{"name": "Heisenberg", "gender": "M"},
		{"name": "Skyler", "gender": "F"},
	}

	// Menampilkan array awal
	fmt.Println("Array Awal:")
	fmt.Println(data)

	// Modifikasi data dalam array
	for i := range data {
		if data[i]["gender"] == "F" {
			data[i]["name"] = "Mrs " + data[i]["name"]
		} else if data[i]["gender"] == "M" {
			data[i]["name"] = "Mr " + data[i]["name"]
		}
	}

	// Menampilkan array setelah dimodifikasi
	fmt.Println("\nArray Setelah Dimodifikasi:")
	fmt.Println(data)
}
