package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	fmt.Println("Pilih operasi aritmatika:")
	fmt.Println("a> penjumlahan (+)")
	fmt.Println("b> pengurangan (-)")
	fmt.Println("c> perkalian (*)")
	fmt.Println("d> pembagian (/)")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Operasi (a/b/c/d): ")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	var result float64

	switch operation {
	case "a":
		fmt.Print("Masukkan angka pertama: ")
		num1, err := readNumber()
		if err != nil {
			panic("Input tidak valid")
		}
		fmt.Print("Masukkan angka kedua: ")
		num2, err := readNumber()
		if err != nil {
			panic("Input tidak valid")
		}
		result = num1 + num2
	case "b":
		fmt.Print("Masukkan angka pertama: ")
		num1, err := readNumber()
		if err != nil {
			panic("Input tidak valid")
		}
		fmt.Print("Masukkan angka kedua: ")
		num2, err := readNumber()
		if err != nil {
			panic("Input tidak valid")
		}
		result = num1 - num2
	case "c":
		fmt.Print("Masukkan angka pertama: ")
		num1, err := readNumber()
		if err != nil {
			panic("Input tidak valid")
		}
		fmt.Print("Masukkan angka kedua: ")
		num2, err := readNumber()
		if err != nil {
			panic("Input tidak valid")
		}
		result = num1 * num2
	case "d":
		fmt.Print("Masukkan angka pertama: ")
		num1, err := readNumber()
		if err != nil {
			panic("Input tidak valid")
		}
		fmt.Print("Masukkan angka kedua: ")
		num2, err := readNumber()
		if err != nil || num2 == 0 {
			panic("Input tidak valid: Pembagian oleh nol atau input tidak valid")
		}
		result = num1 / num2
	default:
		panic("Operasi tidak valid")
	}

	fmt.Printf("Hasil %s adalah %.2f\n", getOperationSymbol(operation), result)
}

func readNumber() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func getOperationSymbol(operation string) string {
	switch operation {
	case "a":
		return "penjumlahan (+)"
	case "b":
		return "pengurangan (-)"
	case "c":
		return "perkalian (*)"
	case "d":
		return "pembagian (/)"
	default:
		return ""
	}
}
