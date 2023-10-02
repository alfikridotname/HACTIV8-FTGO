package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker start %d \n", i)
	<-ch //nungguin pesan channel masuk
	time.Sleep(time.Second)
	fmt.Printf("worker end %d \n", i)
}
func main() {

	var wg sync.WaitGroup
	numWorkers := 10
	ch := make(chan int)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
		ch <- 1 //kirim channel pertama
	}

	wg.Wait()
	close(ch)

	fmt.Println("main end")
}
