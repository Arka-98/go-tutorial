package problems

import (
	"fmt"
	"sync"
	"time"
)

type RWCounter struct {
	rwMu  sync.RWMutex
	count int
}

func (c *RWCounter) incr() {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()

	time.Sleep(500 * time.Millisecond)

	c.count++
}

func (c *RWCounter) get() int {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()

	time.Sleep(500 * time.Millisecond)

	return c.count
}

func RWMutexEx() {
	var wg sync.WaitGroup
	rwCounter := &RWCounter{}
	numGoroutines := 5

	fmt.Println("Before loop")

	for range numGoroutines {
		wg.Go(func() {
			fmt.Println("Get RW count value:", rwCounter.get())
		})
	}

	for i := range numGoroutines {
		wg.Go(func() {
			for range 5 {
				rwCounter.incr()
				fmt.Printf("Incremented goroutine %d\n", i)
			}
		})
	}

	wg.Wait()
}
