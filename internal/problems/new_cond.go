package problems

import (
	"fmt"
	"sync"
	"time"
)

const (
	bufferSize = 3
	count      = 5
)

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	buffer := &buffer{items: make([]int, 0, size)}
	buffer.cond = sync.NewCond(&buffer.mu)

	return buffer
}

func (b *buffer) produce(item int, id int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == bufferSize {
		b.cond.Wait()
	}

	b.items = append(b.items, item)

	fmt.Printf("Thread: %d, Produced: %d\n", id, item)
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for isItemsEmpty(b.items) {
		// fmt.Println("items is empty")
		b.cond.Wait()
	}

	item := b.items[0]

	fmt.Println("Consumed:", item)

	b.items = b.items[1:]

	b.cond.Signal()

	return item
}

func isItemsEmpty(items []int) bool {
	// fmt.Println("ran")

	return len(items) == 0
}

func producer(b *buffer, id int) {
	for i := range count {
		time.Sleep(100 * time.Millisecond)
		b.produce(100 + i, id)
	}
}

func consumer(b *buffer) {
	for range count {
		time.Sleep(800 * time.Millisecond)
		b.consume()
	}
}

func SyncCondEx() {
	var wg sync.WaitGroup
	b := newBuffer(bufferSize)

	wg.Go(func() { producer(b, 1) })
	wg.Go(func() { producer(b, 2) })
	wg.Go(func() { consumer(b) })
	wg.Go(func() { consumer(b) })
	wg.Wait()

	fmt.Println("End")
}
