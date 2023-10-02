// Part 1

// package main

// import (
// 	"fmt"
// )

// func addition(chInput chan [2]float64, chOutput chan float64) {
// 	data := <-chInput
// 	chOutput <- data[0] + data[1]
// }

// func subtraction(chInput chan [2]float64, chOutput chan float64) {
// 	data := <-chInput
// 	chOutput <- data[0] - data[1]
// }

// func main() {
// 	chAddInput := make(chan [2]float64)
// 	chSubInput := make(chan [2]float64)
// 	chOutput := make(chan float64)

// 	go addition(chAddInput, chOutput)
// 	go subtraction(chSubInput, chOutput)

// 	// Kirim data ke goroutine
// 	chAddInput <- [2]float64{5, 3}
// 	chSubInput <- [2]float64{5, 3}

// 	// Menerima dan mencetak hasil
// 	fmt.Println("Hasil Penjumlahan:", <-chOutput)
// 	fmt.Println("Hasil Pengurangan:", <-chOutput)
// }

// Part 2
// package main

// import (
// 	"fmt"
// )

// func addition(chInput chan [2]float64, chOutput chan float64) {
// 	data := <-chInput
// 	chOutput <- data[0] + data[1]
// }

// func subtraction(chInput chan [2]float64, chOutput chan float64) {
// 	data := <-chInput
// 	chOutput <- data[0] - data[1]
// }

// func multiplication(chInput chan [2]float64, chOutput chan float64) {
// 	data := <-chInput
// 	chOutput <- data[0] * data[1]
// }

// func division(chInput chan [2]float64, chOutput chan float64) {
// 	data := <-chInput
// 	if data[1] != 0 {
// 		chOutput <- data[0] / data[1]
// 	} else {
// 		chOutput <- 0.0 // Sebagai contoh sederhana, kita kembalikan 0.0 untuk kasus pembagian dengan 0
// 	}
// }

// func main() {
// 	chAddInput := make(chan [2]float64)
// 	chSubInput := make(chan [2]float64)
// 	chMulInput := make(chan [2]float64)
// 	chDivInput := make(chan [2]float64)
// 	chOutput := make(chan float64, 4) // Menampung hasil dari 4 operasi

// 	go addition(chAddInput, chOutput)
// 	go subtraction(chSubInput, chOutput)
// 	go multiplication(chMulInput, chOutput)
// 	go division(chDivInput, chOutput)

// 	chAddInput <- [2]float64{5, 3}
// 	chSubInput <- [2]float64{5, 3}
// 	chMulInput <- [2]float64{5, 3}
// 	chDivInput <- [2]float64{5, 3}

// 	// Karena kita tidak tahu urutan pasti eksekusi goroutine, hasil mungkin dicetak dalam urutan yang berbeda.
// 	for i := 0; i < 4; i++ {
// 		fmt.Println("Hasil:", <-chOutput)
// 	}
// }

// Part 3

// package main

// import (
// 	"fmt"
// )

// type Result struct {
// 	Operation string
// 	Value     float64
// }

// func addition(chInput chan [2]float64, chOutput chan Result) {
// 	data := <-chInput
// 	chOutput <- Result{"tambah", data[0] + data[1]}
// }

// func subtraction(chInput chan [2]float64, chOutput chan Result) {
// 	data := <-chInput
// 	chOutput <- Result{"kurang", data[0] - data[1]}
// }

// func multiplication(chInput chan [2]float64, chOutput chan Result) {
// 	data := <-chInput
// 	chOutput <- Result{"kali", data[0] * data[1]}
// }

// func division(chInput chan [2]float64, chOutput chan Result) {
// 	data := <-chInput
// 	if data[1] != 0 {
// 		chOutput <- Result{"bagi", data[0] / data[1]}
// 	} else {
// 		chOutput <- Result{"bagi", 0.0} // Sebagai contoh sederhana, kita kembalikan 0.0 untuk kasus pembagian dengan 0
// 	}
// }

// func main() {
// 	chAddInput := make(chan [2]float64)
// 	chSubInput := make(chan [2]float64)
// 	chMulInput := make(chan [2]float64)
// 	chDivInput := make(chan [2]float64)
// 	chOutput := make(chan Result, 4) // Menampung hasil dari 4 operasi

// 	go addition(chAddInput, chOutput)
// 	go subtraction(chSubInput, chOutput)
// 	go multiplication(chMulInput, chOutput)
// 	go division(chDivInput, chOutput)

// 	chAddInput <- [2]float64{5, 3}
// 	chSubInput <- [2]float64{5, 3}
// 	chMulInput <- [2]float64{5, 3}
// 	chDivInput <- [2]float64{5, 3}

// 	// Karena kita tidak tahu urutan pasti eksekusi goroutine, hasil mungkin dicetak dalam urutan yang berbeda.
// 	for i := 0; i < 4; i++ {
// 		result := <-chOutput
// 		fmt.Println("Hasil", result.Operation, ":", result.Value)
// 	}
// }

package main

import (
	"fmt"
)

// A. Menambahkan tipe Result
type Result struct {
	Operation string
	Value     float64
}

func addition(chInput chan [2]float64, chOutput chan Result) { // B. Mengubah tipe parameter chOutput
	data := <-chInput
	// C. Mengirimkan struct Result
	chOutput <- Result{"tambah", data[0] + data[1]}
}

func subtraction(chInput chan [2]float64, chOutput chan Result) { // B. Mengubah tipe parameter chOutput
	data := <-chInput
	// C. Mengirimkan struct Result
	chOutput <- Result{"kurang", data[0] - data[1]}
}

func multiplication(chInput chan [2]float64, chOutput chan Result) { // B. Mengubah tipe parameter chOutput
	data := <-chInput
	// C. Mengirimkan struct Result
	chOutput <- Result{"kali", data[0] * data[1]}
}

func division(chInput chan [2]float64, chOutput chan Result) { // B. Mengubah tipe parameter chOutput
	data := <-chInput
	if data[1] != 0 {
		// C. Mengirimkan struct Result
		chOutput <- Result{"bagi", data[0] / data[1]}
	} else {
		// C. Mengirimkan struct Result
		chOutput <- Result{"bagi", 0.0} // Sebagai contoh sederhana, kita kembalikan 0.0 untuk kasus pembagian dengan 0
	}
}

func main() {
	chAddInput := make(chan [2]float64)
	chSubInput := make(chan [2]float64)
	chMulInput := make(chan [2]float64)
	chDivInput := make(chan [2]float64)
	// Mengubah tipe chOutput menjadi Result
	chOutput := make(chan Result, 4) // Menampung hasil dari 4 operasi

	go addition(chAddInput, chOutput)
	go subtraction(chSubInput, chOutput)
	go multiplication(chMulInput, chOutput)
	go division(chDivInput, chOutput)

	chAddInput <- [2]float64{5, 3}
	chSubInput <- [2]float64{5, 3}
	chMulInput <- [2]float64{5, 3}
	chDivInput <- [2]float64{5, 3}

	// D. Modifikasi cara mencetak hasil
	for i := 0; i < 4; i++ {
		result := <-chOutput
		fmt.Println("Hasil", result.Operation, ":", result.Value)
	}
}
