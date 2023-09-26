package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("\nAsynchronous Fibonacci Calculation:")
	var wg sync.WaitGroup
	results := make([]int, 10)

	for n := 1; n <= 10; n++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			result := fibonacci(n)
			results[n-1] = result
			fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", n, result)
		}(n)
	}

	wg.Wait()
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
