package problems

import (
	"fmt"
	"sync"
	"time"
)

func WorkerWg(id int, wg *sync.WaitGroup, tasks <-chan int, results chan<- int) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Goroutine %d. Task: %d\n", id, task)

		// Processing
		time.Sleep(time.Second)

		results <- task * 2

		fmt.Printf("Goroutine %d. Finished task %d\n", id, task)
	}

	fmt.Printf("Goroutine %d complete\n", id)
}

func ProcessWg() {
	var wg sync.WaitGroup

	numTasks := 10
	numWorkers := 3
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	// Initiate workers
	for i := range numWorkers {
		wg.Add(1)

		go WorkerWg(i+1, &wg, tasks, results)
	}

	// Send tasks
	for i := range numTasks {
		tasks <- i + 1
	}

	close(tasks)

	fmt.Println("Tasks sent")

	go func() {
		wg.Wait()

		close(results)
	}()

	// Wait for results
	for result := range results {
		fmt.Println("Received:", result)
	}

	fmt.Println("End")
}
