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

	time.Sleep(2000 * time.Millisecond)

	return c.count
}

func RWMutexEx() {
	var (
		wg        sync.WaitGroup
		rwCounter RWCounter
	)

	numGoroutines := 5

	fmt.Println("Before loop")

	for range numGoroutines {
		wg.Go(func() {
			fmt.Printf("Get RW count value: %d %v\n", rwCounter.get(), time.Now())
		})
	}

	for i := range numGoroutines {
		wg.Go(func() {
			for range 1 {
				rwCounter.incr()
				fmt.Printf("Incremented goroutine %d %v\n", i, time.Now())
			}
		})
	}

	wg.Wait()
}
