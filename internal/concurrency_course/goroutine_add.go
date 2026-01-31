package concurrencycourse

import (
	"fmt"
	"sync"
)

var sum int

// var mu sync.Mutex

func Add(val int) {
	sum += val
}

func GoroutineAddTest() {
	var wg sync.WaitGroup

	for range 5 {
		wg.Go(func() {
			Add(1)
		})
	}

	wg.Wait()

	fmt.Println("Sum:", sum)
}
