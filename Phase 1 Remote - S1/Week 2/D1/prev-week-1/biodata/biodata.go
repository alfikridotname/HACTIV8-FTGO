package main

import (
	"fmt"
	"os"
)

// Struct untuk merepresentasikan data teman-teman kalian
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	// Membuat map untuk menyimpan data teman-teman berdasarkan nomor absen
	temanTeman := map[int]Teman{
		1: {"Teman1", "Alamat1", "Pekerjaan1", "Alasan1"},
		2: {"Teman2", "Alamat2", "Pekerjaan2", "Alasan2"},
		3: {"Teman3", "Alamat3", "Pekerjaan3", "Alasan3"},
		// Tambahkan data teman-teman lainnya di sini
	}

	// Mendapatkan nomor absen dari argumen CLI
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Cara pakai: go run biodata.go <nomor_absen>")
		return
	}

	// Mengkonversi nomor absen dari string ke integer
	nomorAbsen := args[1]
	var nomorAbsenInt int
	_, err := fmt.Sscanf(nomorAbsen, "%d", &nomorAbsenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	// Mencari data teman berdasarkan nomor absen
	teman, found := temanTeman[nomorAbsenInt]
	if !found {
		fmt.Println("Teman dengan nomor absen", nomorAbsenInt, "tidak ditemukan")
		return
	}

	// Menampilkan data teman
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}
