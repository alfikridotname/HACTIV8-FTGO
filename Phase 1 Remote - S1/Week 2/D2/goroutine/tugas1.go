package main

import (
	"fmt"
	"sync"
	"time"
)

type LogMessage struct {
	Level   string
	Message string
}

func printLogMessage(m LogMessage) {
	fmt.Println(m.Level, m.Message)
}

func main() {
	var wg sync.WaitGroup

	LogMessages := []LogMessage{
		{Level: "[INFO]", Message: "User 123 logged in"},
		{Level: "[WARN]", Message: "Memory usage is high"},
		{Level: "[ERROR]", Message: "Failed to fetch data from API"},
	}

	start := time.Now()
	fmt.Printf("Start ===============> %d\n", time.Since(start).Milliseconds())
	for i, msg := range LogMessages {
		wg.Add(1)
		go func(m LogMessage, i time.Duration) {
			defer wg.Done()
			time.Sleep(i * time.Second)
			printLogMessage(m)
		}(msg, time.Duration(i+1))
	}

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("[INFO] - User 123 logged in")
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("[WARN] - Memory usage is high")
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("[ERROR] - Failed to fetch data from API")
	// }()

	fmt.Println("Application continue after logging...")

	wg.Wait()

	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Took ===============> %d\n", elapsed)

	fmt.Println("Main application finished")
}
