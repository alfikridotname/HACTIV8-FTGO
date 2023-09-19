package main

import "fmt"

func main() {
	// Menghitung rata rata dari data 1,2,3,4,3,2,1,2,3,4,3,2,1
	data := []int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1}
	var total int
	for _, value := range data {
		total += value
	}
	average := float64(total) / float64(len(data))
	// pembulatan ke atas dan 2 angka di belakang koma
	fmt.Printf("Rata-rata: %.2f\n", average)
}
