package main

import (
	"fmt"
	"ngc_lima/jawaban"
)

func main() {
	fmt.Println("Pilih soal: ")
	fmt.Println("1. Struct & method 1")
	fmt.Println("2. Struct & method 2")
	fmt.Println("3. Struct Hero 1")
	fmt.Println("4. Struct Hero 2")

	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		jawaban.StructMethodSatu()
	case 2:
		jawaban.StructMethodDua()
	case 3:
		jawaban.StructHeroSatu()
	case 4:
		jawaban.StructHeroDua()
	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}
