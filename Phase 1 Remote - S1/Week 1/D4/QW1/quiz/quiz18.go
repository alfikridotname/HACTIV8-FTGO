package quiz

import "fmt"

func Quiz18() {
	var n int

	fmt.Print("Masukkan nilai n (bilangan ganjil): ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("n harus merupakan bilangan ganjil.")
		return
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
