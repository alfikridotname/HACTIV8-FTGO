package quiz

import "fmt"

func Quiz2() {
	// Membuat map untuk menyimpan harga buah
	hargaBuah := make(map[string]float64)

	// Mengisi map dengan data buah dan harga per kilogram
	hargaBuah["Apel"] = 12.5
	hargaBuah["Mangga"] = 15.0
	hargaBuah["Pisang"] = 8.75

	// Menampilkan data dalam map
	fmt.Println("Data Harga Buah:")
	for buah, harga := range hargaBuah {
		fmt.Printf("%s: %.2f per kilogram\n", buah, harga)
	}
}
