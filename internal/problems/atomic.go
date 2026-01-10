package problems

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int32
}

func (c *AtomicCounter) incrBy(val int32) {
	atomic.AddInt32(&c.count, val)
}

func (c *AtomicCounter) get() int32 {
	return atomic.LoadInt32(&c.count)
}

func AtomicExample() {
	var wg sync.WaitGroup

	counter := &AtomicCounter{}
	// value := 0
	numGoroutines := 10

	for range numGoroutines {
		wg.Go(func() {
			for range 1000 {
				counter.incrBy(1)
				// value++
			}
		})
	}

	wg.Wait()
	fmt.Println("Total count:", counter.get())
	// fmt.Println("Total count:", value)
}
