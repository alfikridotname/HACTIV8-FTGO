package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Minta pengguna untuk memasukkan angka
	fmt.Print("Masukkan angka: ")
	input, err := getInput()
	if err != nil {
		fmt.Println("Input tidak valid. Masukkan angka positif.")
		return
	}

	// Cek apakah angka tersebut genap atau ganjil
	if input%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// Cetak semua angka genap dari 1 hingga input
	fmt.Println("Angka genap:")
	for i := 1; i <= input; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// Fungsi untuk mendapatkan input dari pengguna dan mengembalikannya sebagai integer
func getInput() (int, error) {
	var inputStr string
	fmt.Scanln(&inputStr)
	input, err := strconv.Atoi(inputStr)
	return input, err
}
