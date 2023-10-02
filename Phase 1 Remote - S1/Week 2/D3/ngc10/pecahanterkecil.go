package main

import "fmt"

// Fungsi untuk mencari FPB (Faktor Persekutuan Terbesar)
func fpb(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Masukkan pecahan
	numerator := 3
	denominator := 9

	// Cari FPB dari pembilang dan penyebut
	gcd := fpb(numerator, denominator)

	// Sederhanakan pecahan
	simplifiedNumerator := numerator / gcd
	simplifiedDenominator := denominator / gcd

	fmt.Printf("%d/%d -> %d/%d\n", numerator, denominator, simplifiedNumerator, simplifiedDenominator)
}
