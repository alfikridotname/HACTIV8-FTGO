package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var pilihan uint8
	fmt.Println("Masukkan Pilihan Soal Anda (1-9): ")
	fmt.Println("1. Looping 1")
	fmt.Println("2. Looping 2")
	fmt.Println("3. Palindrome")
	fmt.Println("4. XOXO")
	fmt.Println("5. XOXO Order")
	fmt.Println("6. Asterik Level 1")
	fmt.Println("7. Asterik Level 2")
	fmt.Println("8. Asterik Level 3")
	fmt.Println("9. Asterik Level 4")

	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		looping1()
	case 2:
		looping2()
	case 3:
		palindrome()
	case 4:
		xoxo()
	case 5:
		xoxoOrder()
	case 6:
		asterikLevel1()
	case 7:
		asterikLevel2()
	case 8:
		asterikLevel3()
	case 9:
		asterikLevel4()
	}
}

func looping1() {
	// Membuat slice yang berisi data map orang-orang
	people := []map[string]any{
		{
			"name": "Hank",
			"Age":  50,
			"Job":  "Polisi",
		},
		{
			"name": "Heisenberg",
			"Age":  52,
			"Job":  "Ilmuwan",
		},
		{
			"name": "Skyler",
			"Age":  48,
			"Job":  "Akuntan",
		},
	}

	// Melakukan looping dan menampilkan data tiap orang dengan format yang diinginkan
	for _, person := range people {
		name := person["name"]
		age := person["Age"]
		job := person["Job"]

		fmt.Printf("Hi, Perkenalkan, Nama saya %s, umur saya %d, dan saya bekerja sebagai %s\n", name, age, job)
	}
}

func looping2() {
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	// Menghitung rata-rata, jumlah, dan median untuk slice1
	avg1 := calculateAverage(slice1)
	sum1 := calculateSum(slice1)
	median1 := calculateMedian(slice1)

	// Menghitung rata-rata, jumlah, dan median untuk slice2
	avg2 := calculateAverage(slice2)
	sum2 := calculateSum(slice2)
	median2 := calculateMedian(slice2)

	// Menampilkan hasil perhitungan
	fmt.Printf("Slice 1:\n")
	fmt.Printf("Rata-rata: %.2f\n", avg1)
	fmt.Printf("Jumlah: %.2f\n", sum1)
	fmt.Printf("Median: %.2f\n\n", median1)

	fmt.Printf("Slice 2:\n")
	fmt.Printf("Rata-rata: %.2f\n", avg2)
	fmt.Printf("Jumlah: %.2f\n", sum2)
	fmt.Printf("Median: %.2f\n", median2)

	// kembali ke menu utama
	main()
}

func calculateAverage(slice []float64) float64 {
	sum := calculateSum(slice)
	return sum / float64(len(slice))
}

func calculateSum(slice []float64) float64 {
	sum := 0.0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func calculateMedian(slice []float64) float64 {
	sort.Float64s(slice)
	length := len(slice)

	if length%2 == 0 {
		// Jika jumlah elemen genap, ambil rata-rata dari dua nilai tengah
		mid1 := slice[length/2-1]
		mid2 := slice[length/2]
		return (mid1 + mid2) / 2
	} else {
		// Jika jumlah elemen ganjil, ambil nilai tengah
		return slice[length/2]
	}
}

func palindrome() {
	var kata string
	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)

	// Menghapus spasi dan mengubah kata menjadi huruf kecil
	kata = strings.ToLower(strings.ReplaceAll(kata, " ", ""))

	// Memeriksa apakah kata adalah palindrome
	if isPalindrome(kata) {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}
}

func isPalindrome(kata string) bool {
	// Menghitung panjang kata
	panjang := len(kata)

	// Membandingkan setiap karakter dari awal dan akhir kata
	for i := 0; i < panjang/2; i++ {
		if kata[i] != kata[panjang-i-1] {
			return false
		}
	}
	return true
}

func xoxo() {
	var kata string
	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)

	// Menghitung jumlah karakter 'x' dan 'o' dalam kata
	countX := strings.Count(kata, "x")
	countO := strings.Count(kata, "o")

	// Membandingkan jumlah karakter 'x' dan 'o'
	sama := countX == countO

	fmt.Println(sama)
}

func xoxoOrder() {
	var kata string
	fmt.Print("Masukkan angka-angka (pisahkan dengan spasi): ")
	fmt.Scan(&kata)

	// Mengubah string menjadi slice of int
	angkaStr := strings.Fields(kata)
	angka := make([]int, len(angkaStr))

	for i, str := range angkaStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Input tidak valid.")
			return
		}
		angka[i] = num
	}

	// Mengurutkan slice dari besar ke kecil menggunakan Bubble Sort
	n := len(angka)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if angka[j] < angka[j+1] {
				// Tukar angka[j] dan angka[j+1]
				angka[j], angka[j+1] = angka[j+1], angka[j]
			}
		}
	}

	// Menampilkan slice yang sudah diurutkan
	fmt.Println("Slice yang sudah diurutkan dari besar ke kecil:")
	// Tampilkan semua angka
	for _, num := range angka {
		fmt.Printf("%d ", num)
	}
}

func asterikLevel1() {
	var rows1 int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&rows1)

	for i := 0; i < rows1; i++ {
		fmt.Print("*")
		fmt.Println()
	}
}

func asterikLevel2() {
	var rows2 int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&rows2)

	for i := 0; i < rows2; i++ {
		for j := 0; j < rows2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func asterikLevel3() {
	var rows3 int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&rows3)

	for i := 1; i <= rows3; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func asterikLevel4() {
	var rows4 int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&rows4)

	for i := rows4; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
