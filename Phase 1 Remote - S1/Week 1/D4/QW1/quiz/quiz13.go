package quiz

import "fmt"

func Quiz13() {
	hasil := jumlahBilangan(5, 7, 10, 2)

	fmt.Printf("Hasil penjumlahan bilangan adalah %d\n", hasil)
}

// Fungsi variadic untuk menghitung jumlah bilangan bulat
func jumlahBilangan(bilangan ...int) int {
	total := 0
	for _, nilai := range bilangan {
		total += nilai
	}
	return total
}
