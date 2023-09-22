package quiz

import "fmt"

// Buat struct untuk merepresentasikan informasi buku
type Buku struct {
	Judul   string
	Penulis string
	Harga   float64
}

func Quiz15() {
	// Membuat beberapa contoh buku
	buku1 := Buku{"Harry Potter and the Sorcerer's Stone", "J.K. Rowling", 20.99}
	buku2 := Buku{"The Hobbit", "J.R.R. Tolkien", 15.99}
	buku3 := Buku{"To Kill a Mockingbird", "Harper Lee", 12.49}

	// Membuat slice dari struct buku
	daftarBuku := []Buku{buku1, buku2, buku3}

	// Menghitung total harga buku
	totalHarga := TotalHarga(daftarBuku)

	fmt.Printf("Total harga semua buku adalah $%.2f\n", totalHarga)
}

// Fungsi untuk menghitung total harga dari slice buku
func TotalHarga(buku []Buku) float64 {
	total := 0.0
	for _, buku := range buku {
		total += buku.Harga
	}
	return total
}
