package quiz

import "fmt"

func Quiz10() {
	tinggi := 5

	for baris := 0; baris < tinggi; baris++ {
		nilai := 1 // Nilai awal pada setiap baris

		// Mencetak spasi untuk pembentukan segitiga
		for spasi := 0; spasi < tinggi-baris; spasi++ {
			fmt.Print("  ")
		}

		// Mencetak nilai pada setiap baris
		for kolom := 0; kolom <= baris; kolom++ {
			fmt.Printf("%4d", nilai)
			nilai = nilai * (baris - kolom) / (kolom + 1)
		}

		fmt.Println() // Pindah ke baris berikutnya setelah mencetak satu baris
	}
}
