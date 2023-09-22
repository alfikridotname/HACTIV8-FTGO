package quiz

import "fmt"

func Quiz12() {
	bilangan1 := 5
	bilangan2 := 7

	hasil := jumlah(bilangan1, bilangan2)

	fmt.Printf("Hasil penjumlahan %d dan %d adalah %d\n", bilangan1, bilangan2, hasil)
}

// Fungsi untuk menghitung jumlah dua bilangan bulat
func jumlah(a, b int) int {
	return a + b
}
