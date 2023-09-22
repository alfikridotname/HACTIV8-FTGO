package quiz

import "fmt"

func Quiz6() {
	var N int

	fmt.Print("Masukkan jumlah suku pertama yang ingin dicetak: ")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("Jumlah suku harus lebih besar dari 0.")
		return
	}

	// Menginisialisasi suku pertama
	suku := 1

	fmt.Println("Deret geometri:")
	for i := 0; i < N; i++ {
		fmt.Print(suku, " ")

		// Menggunakan rasio 2 untuk menghasilkan suku berikutnya
		suku *= 2
	}
	fmt.Println()
}
