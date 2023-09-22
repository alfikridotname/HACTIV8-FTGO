package quiz

import "fmt"

func Quiz11() {
	bilangan1 := 5
	bilangan2 := 7

	hasil := tambah(bilangan1, bilangan2)

	fmt.Printf("Hasil penjumlahan %d dan %d adalah %d\n", bilangan1, bilangan2, hasil)
}

// Fungsi untuk menghitung jumlah dua bilangan
func tambah(a, b int) int {
	return a + b
}
