package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	scheduleTask := func(taskID int) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			task(taskID)
		}()
	}

	// Number of tasks
	numTasks := 3

	fmt.Println("Main application is continuing...")

	for i := 1; i <= numTasks; i++ {
		scheduleTask(i)
	}

	wg.Wait()

	fmt.Println("All tasks completed.")
}

func task(taskID int) {
	fmt.Printf("Task %d being executed\n", taskID)
	time.Sleep(time.Duration(taskID) * time.Second)
}
