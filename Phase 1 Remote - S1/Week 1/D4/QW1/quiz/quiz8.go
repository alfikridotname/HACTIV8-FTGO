package quiz

import "fmt"

func Quiz8() {
	var N int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("N harus lebih besar dari 0.")
		return
	}

	fmt.Println("Hasil:")
	for i := 1; i <= N; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
