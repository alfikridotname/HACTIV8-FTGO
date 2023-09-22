package quiz

import "fmt"

func Quiz3() {
	var bilangan int

	fmt.Print("Masukkan suatu bilangan: ")
	fmt.Scan(&bilangan)

	// Menggunakan struktur pengkondisian if else
	if bilangan%2 == 0 {
		fmt.Printf("%d adalah bilangan genap\n", bilangan)
	} else {
		fmt.Printf("%d adalah bilangan ganjil\n", bilangan)
	}
}
