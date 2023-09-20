package jawaban

import "fmt"

type personMethodSatu struct {
	Name string
	Age  int
	Job  string
}

// Metode GetInfo untuk menampilkan informasi Person
func (p *personMethodSatu) getInfo() {
	fmt.Printf("Name: %s\nAge: %d\nJob: %s\n", p.Name, p.Age, p.Job)
}

// Metode AddYear untuk menambahkan 1 tahun pada usia dan mengubah pekerjaan jika diperlukan
func (p *personMethodSatu) AddYear() {
	p.Age++ // Tambahkan 1 tahun pada usia

	if p.Age >= 50 {
		p.Job = "Retired"
	}
}

func StructMethodSatu() {
	p := &personMethodSatu{
		Name: "Bambang",
		Age:  45,
		Job:  "Gambler",
	}

	fmt.Println("Informasi awal:")
	p.getInfo()

	// Tambahkan usia sebanyak 5 kali
	for i := 0; i < 5; i++ {
		p.AddYear()
	}

	fmt.Println("\nInformasi setelah menambahkan usia:")
	p.getInfo()
}
