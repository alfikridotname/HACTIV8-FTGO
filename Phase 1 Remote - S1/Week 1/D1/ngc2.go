package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var pilihan uint8
	fmt.Print("Masukkan Pilihan Soal Anda (1-2): ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		// Conditional 1
		conditional1()
	} else {
		// Conditional 2
		conditional2()
	}
}

func conditional1() {
	rand.Seed(time.Now().UnixNano())

	var name string
	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&name)

	randomNumber := rand.Intn(100) + 1

	fmt.Printf("Selamat %s, ", name)

	switch {
	case randomNumber > 80:
		fmt.Println("anda sangat beruntung")
	case randomNumber <= 80 && randomNumber > 60:
		fmt.Println("anda beruntung")
	case randomNumber <= 60 && randomNumber > 40:
		fmt.Println("anda kurang beruntung")
	default:
		fmt.Println("anda sial")
	}

	os.Exit(0)
}

func conditional2() {
	fmt.Println("Selamat datang di Aplikasi Verifikasi Umur")
	fmt.Print("Masukkan Nama Anda: ")

	reader := bufio.NewReader(os.Stdin)
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Umur Anda: ")

	umurStr, _ := reader.ReadString('\n')
	umurStr = strings.TrimSpace(umurStr)

	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		fmt.Println("PESAN ERROR: Umur harus berupa angka.")
		return
	}

	if umur < 0 || umur > 100 {
		fmt.Println("PESAN ERROR: Umur harus berada dalam rentang 0 hingga 100 tahun.")
		return
	}

	if umur > 18 {
		fmt.Printf("Halo, %s! Silahkan masuk.\n", nama)
	} else {
		fmt.Printf("Halo, %s! Dilarang masuk, maksimal umur 19.\n", nama)
	}
}
