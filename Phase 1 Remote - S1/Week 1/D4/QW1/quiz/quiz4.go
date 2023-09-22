package quiz

import "fmt"

func Quiz4() {
	var angkaBulan int

	fmt.Print("Masukkan angka bulan (1-12): ")
	fmt.Scan(&angkaBulan)

	var namaBulan string

	switch angkaBulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		fmt.Println("Input tidak valid. Masukkan angka bulan antara 1-12.")
		return
	}

	fmt.Printf("Bulan %d adalah %s\n", angkaBulan, namaBulan)
}
