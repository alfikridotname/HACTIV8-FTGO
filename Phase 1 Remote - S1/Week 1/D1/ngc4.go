package main

import (
	"fmt"
	"strings"
)

func main() {
	// AlayGen
	a := AlayGen1("Hello", "World", "zzz")
	fmt.Println(a)

	// Fibo
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibo(i))
	}
	fmt.Println()
}

func AlayGen(kata ...string) {
	panjang := len(kata)
	for i := 0; i < panjang; i++ {
		for j := 0; j < len(kata[i]); j++ {
			tes := konversi(string(kata[i][j]))
			fmt.Print(tes)
		}
		fmt.Print(" ")
	}
}

func konversi(huruf string) string {
	var hurufKonversi string
	if huruf == "a" {
		hurufKonversi += "4"
	} else if huruf == "e" {
		hurufKonversi += "3"
	} else if huruf == "i" {
		hurufKonversi += "1"
	} else if huruf == "l" {
		hurufKonversi += "1"
	} else if huruf == "n" {
		hurufKonversi += "N"
	} else if huruf == "s" {
		hurufKonversi += "5"
	} else if huruf == "x" {
		hurufKonversi += "*"
	} else {
		hurufKonversi += huruf
	}

	return hurufKonversi
}

// Fungsi variadik untuk menggabungkan string-string 'alay' dari mas safran
func AlayGen1(kata ...string) string {
	var hasil strings.Builder
	for i, s := range kata {
		if i > 0 {
			hasil.WriteString(" ")
		}
		hasil.WriteString(konversi(s))
	}
	return hasil.String()
}

// Fungsi untuk fibonaci
func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}
