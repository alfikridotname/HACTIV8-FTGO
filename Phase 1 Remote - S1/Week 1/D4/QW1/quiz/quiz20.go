package quiz

import "fmt"

func Quiz20() {
	var n int

	fmt.Print("Masukkan jumlah siswa: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah siswa harus lebih dari 0.")
		return
	}

	for i := 1; i <= n; i++ {
		var nilai int
		var predikat string

		fmt.Printf("Masukkan nilai siswa %d: ", i)
		fmt.Scan(&nilai)

		if nilai < 50 {
			predikat = "D"
		} else if nilai >= 50 && nilai <= 69 {
			predikat = "C"
		} else if nilai >= 70 && nilai <= 84 {
			predikat = "B"
		} else if nilai >= 85 {
			predikat = "A"
		}

		fmt.Printf("Siswa %d mendapatkan predikat %s\n", i, predikat)
	}
}
