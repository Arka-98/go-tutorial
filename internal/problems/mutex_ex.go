package problems

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) incr() {
	defer c.mu.Unlock()
	c.mu.Lock()

	c.count++
}

func (c *Counter) get() int {
	defer c.mu.Unlock()
	c.mu.Lock()

	return c.count
}

func MutexEx() {
	var wg sync.WaitGroup

	counter := &Counter{}
	numGoroutines := 10

	for range numGoroutines {
		fmt.Println("ran")

		wg.Go(func() {
			for range 1000 {
				counter.incr()
			}
		})
	}

	wg.Wait()
	fmt.Println("Total count:", counter.get())
}
