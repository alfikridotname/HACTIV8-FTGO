package jawaban

import (
	"fmt"
	"math/rand"
)

// Definisikan struct Weapon
type WeaponHeroDua struct {
	Attack int
}

// Definisikan struct Hero
type HeroDua struct {
	Name           string
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealthPoint    int
	Weapon         WeaponHeroDua // Properti Weapon dalam struct Hero
}

// Metode CountDamage untuk menghitung total damage
func (h *HeroDua) CountDamage() int {
	totalDamage := h.BaseAttack + h.Weapon.Attack
	totalDamage -= h.Defence

	// Tentukan peluang critical damage 50:50
	if rand.Intn(2) == 1 {
		totalDamage += h.CriticalDamage
	}

	return totalDamage
}

// Metode isAttackedBy untuk menghitung total damage yang diterima dan pengurangan health point
func (h *HeroDua) isAttackedBy(attacker HeroDua) {
	totalDamageReceived := attacker.CountDamage()

	if totalDamageReceived < 0 {
		totalDamageReceived = 0
	}

	h.HealthPoint -= totalDamageReceived
}

// Fungsi Battle untuk simulasi pertarungan antara dua Hero
func Battle(attacker HeroDua, defender HeroDua) {
	fmt.Printf("%s menyerang %s!\n", attacker.Name, defender.Name)

	defender.isAttackedBy(attacker)

	fmt.Printf("%s menerima total damage sebesar %d dan memiliki health point sisa %d\n", defender.Name, attacker.CountDamage(), defender.HealthPoint)
}

func StructHeroDua() {
	// Buat dua instance Hero untuk simulasi pertarungan
	hero1 := HeroDua{
		Name:           "Hero1",
		BaseAttack:     30,
		Defence:        10,
		CriticalDamage: 20,
		HealthPoint:    100,
		Weapon: WeaponHeroDua{
			Attack: 15,
		},
	}

	hero2 := HeroDua{
		Name:           "Hero2",
		BaseAttack:     25,
		Defence:        15,
		CriticalDamage: 15,
		HealthPoint:    120,
		Weapon: WeaponHeroDua{
			Attack: 20,
		},
	}

	// Mulai pertarungan
	Battle(hero1, hero2)
}
