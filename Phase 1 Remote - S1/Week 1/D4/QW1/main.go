package main

import (
	"fmt"
	"qw1/quiz"
)

func main() {
	fmt.Println("Pilih kuis yang ingin dijalankan:")

	for i := 1; i <= 20; i++ {
		fmt.Printf("%d. Kuis %d\n", i, i)
	}

	var kuis int
	fmt.Scanln(&kuis)

	// switch 20 case
	switch kuis {
	case 1:
		quiz.Quiz1()
	case 2:
		quiz.Quiz2()
	case 3:
		quiz.Quiz3()
	case 4:
		quiz.Quiz4()
	case 5:
		quiz.Quiz5()
	case 6:
		quiz.Quiz6()
	case 7:
		quiz.Quiz7()
	case 8:
		quiz.Quiz8()
	case 9:
		quiz.Quiz9()
	case 10:
		quiz.Quiz10()
	case 11:
		quiz.Quiz11()
	case 12:
		quiz.Quiz12()
	case 13:
		quiz.Quiz13()
	case 14:
		quiz.Quiz14()
	case 15:
		quiz.Quiz15()
	case 16:
		quiz.Quiz16()
	case 17:
		quiz.Quiz17()
	case 18:
		quiz.Quiz18()
	case 19:
		quiz.Quiz19()
	case 20:
		quiz.Quiz20()
	default:
		fmt.Println("Quiz tidak ditemukan")
	}
}
