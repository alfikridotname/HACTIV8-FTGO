package quiz

import "fmt"

func Quiz14() {
	hasil := nilaiMaksimum(5, 7, 10, 2, 8, 3)

	fmt.Printf("Nilai maksimum dari bilangan adalah %d\n", hasil)
}

// Fungsi variadic untuk mencari nilai maksimum dari bilangan bulat
func nilaiMaksimum(bilangan ...int) int {
	if len(bilangan) == 0 {
		return 0 // Kembalikan 0 jika tidak ada bilangan yang diberikan
	}

	maksimum := bilangan[0] // Inisialisasi dengan nilai pertama

	for _, nilai := range bilangan {
		if nilai > maksimum {
			maksimum = nilai
		}
	}

	return maksimum
}
