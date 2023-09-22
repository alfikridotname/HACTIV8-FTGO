package quiz

import "fmt"

func Quiz19() {
	kalimat, jumlahKata := buildSentence("Ini", "adalah", "contoh", "fungsi", "variadic")
	fmt.Printf("Kalimat: %s\n", kalimat)
	fmt.Printf("Jumlah Kata: %d\n", jumlahKata)
}

// Fungsi variadic untuk menggabungkan kata-kata menjadi kalimat
func buildSentence(kata ...string) (kalimat string, jumlahKata int) {
	for i, k := range kata {
		kalimat += k
		jumlahKata++
		if i < len(kata)-1 {
			kalimat += " "
		}
	}
	return kalimat, jumlahKata
}
