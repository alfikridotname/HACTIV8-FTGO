package quiz

import "fmt"

func Quiz9() {
	var N int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("N harus lebih besar dari 0.")
		return
	}

	fmt.Println("Hasil:")
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			// Cetak '*' di tepi luar atau jika i atau j adalah batas kotak
			// Cetak spasi di dalam kotak
			if i == 1 || i == N || j == 1 || j == N {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

			// Tambahkan spasi antara setiap karakter
			fmt.Print(" ")
		}
		fmt.Println() // Pindah ke baris berikutnya setelah mencetak satu baris
	}
}
