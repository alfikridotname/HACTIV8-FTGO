package main

import "fmt"

func main() {
	chInput := make(chan [2]float64)
	chOutput := make(chan float64, 2)

	go func() {
		data := <-chInput
		chOutput <- data[0] + data[1]
		data = <-chInput
		chOutput <- data[0] - data[1]
	}()

	chInput <- [2]float64{5, 3}
	chInput <- [2]float64{6, 1}

	for i := 0; i < 2; i++ {
		select {
		case result := <-chOutput:
			hasil := fmt.Sprintf("Hasil: %f\n", result)
			fmt.Println(hasil)
		}
	}
}
