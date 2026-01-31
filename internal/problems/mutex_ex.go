package problems

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) incr() {
	c.mu.Lock()
	defer c.mu.Unlock()

	time.Sleep(time.Millisecond)

	c.count++
}

func (c *Counter) get() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	time.Sleep(500 * time.Millisecond)

	return c.count
}

func MutexEx() {
	var wg sync.WaitGroup

	counter := &Counter{}
	numGoroutines := 10
	start := time.Now()

	// for i := range numGoroutines {
	// 	fmt.Println("ran goroutine", i)
	// 	wg.Go(func() {
	// 		for range 1000 {
	// 			counter.incr()
	// 		}
	// 	})
	// }

	wg.Wait()
	fmt.Printf("Total count: %d, Elapsed: %v", counter.get(), time.Since(start))

	for i := range numGoroutines {
		fmt.Println("ran goroutine", i)
		wg.Go(func() {
			fmt.Println("Get count value:", counter.get())
		})
	}

	wg.Wait()
	fmt.Println("End")
}
