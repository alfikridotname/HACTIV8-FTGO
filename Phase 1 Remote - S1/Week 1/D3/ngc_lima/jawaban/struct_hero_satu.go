package jawaban

import (
	"fmt"
	"math/rand"
	"time"
)

// Definisikan struct Weapon
type WeaponHeroSatu struct {
	Attack int
}

// Definisikan struct Hero
type HeroSatu struct {
	Name           string
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealthPoint    int
	Weapon         WeaponHeroSatu // Properti Weapon dalam struct Hero
}

// Metode CountDamage untuk menghitung total damage
func (h *HeroSatu) CountDamage() int {
	// Hitung total damage
	totalDamage := h.BaseAttack + h.Weapon.Attack

	// Tentukan peluang critical damage 50:50
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(2)

	if chance == 1 {
		totalDamage += h.CriticalDamage
	}

	return totalDamage
}

func StructHeroSatu() {
	// Buat sebuah instance Hero
	hero := HeroSatu{
		Name:           "Hero1",
		BaseAttack:     30,
		Defence:        10,
		CriticalDamage: 20,
		HealthPoint:    100,
		Weapon: WeaponHeroSatu{
			Attack: 15,
		},
	}

	// Hitung total damage saat hero menyerang
	totalDamage := hero.CountDamage()

	fmt.Printf("Total Damage saat serangan oleh %s: %d\n", hero.Name, totalDamage)
}
