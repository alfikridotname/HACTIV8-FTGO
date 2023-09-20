package jawaban

import "fmt"

type personMethodDua struct {
	Name string
	Age  int
	Job  string
}

// Metode GetInfo untuk menampilkan informasi Person
func (p *personMethodDua) GetInfo() {
	fmt.Printf("Name: %s\nAge: %d\nJob: %s\n", p.Name, p.Age, p.Job)
}

func StructMethodDua() {
	// Buat beberapa instance Person
	persons := []personMethodDua{
		{Name: "Bambang", Age: 20, Job: "Gambler"},
		{Name: "Alice", Age: 30, Job: "Engineer"},
		{Name: "Charlie", Age: 25, Job: "Teacher"},
	}

	// Lakukan pengulangan pada slice persons dan tampilkan informasi dengan GetInfo
	for _, person := range persons {
		fmt.Println("Informasi:")
		person.GetInfo()
		fmt.Println() // Tambahkan baris kosong antara setiap informasi
	}
}
