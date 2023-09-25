package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Mengecek jumlah argumen yang benar
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <operation> <operand1> <operand2>")
		return
	}

	// Mendapatkan argumen dari command line
	operation := os.Args[1]
	operand1Str := os.Args[2]
	operand2Str := os.Args[3]

	// Mengkonversi operand1 dan operand2 ke float64
	operand1, err1 := strconv.ParseFloat(operand1Str, 64)
	operand2, err2 := strconv.ParseFloat(operand2Str, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Operand harus berupa angka")
		return
	}

	// Melakukan operasi sesuai dengan perintah
	result := 0.0
	switch operation {
	case "add":
		result = operand1 + operand2
	case "sub":
		result = operand1 - operand2
	case "mul":
		result = operand1 * operand2
	case "div":
		if operand2 == 0 {
			fmt.Println("Pembagian dengan nol tidak diizinkan")
			return
		}
		result = operand1 / operand2
	default:
		fmt.Println("Operasi yang diizinkan: addition, subtraction, multiplication, division")
		return
	}

	// Menampilkan hasil
	fmt.Printf("Hasil %s %.2f dan %.2f adalah %.2f\n", operation, operand1, operand2, result)
}
