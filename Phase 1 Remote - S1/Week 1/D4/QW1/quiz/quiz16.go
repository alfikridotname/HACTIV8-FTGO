package quiz

import "fmt"

func Quiz16() {
	var performa string
	var gaji float64
	var bonus float64

	fmt.Print("Masukkan performa karyawan (A/B/C/D): ")
	fmt.Scan(&performa)

	fmt.Print("Masukkan gaji karyawan: ")
	fmt.Scan(&gaji)

	switch performa {
	case "A":
		bonus = 0.20 * gaji
	case "B":
		bonus = 0.10 * gaji
	case "C":
		bonus = 0.05 * gaji
	case "D":
		bonus = 0
	default:
		fmt.Println("Performa tidak valid.")
		return
	}

	fmt.Printf("Bonus untuk performa %s adalah $%.2f\n", performa, bonus)
}
