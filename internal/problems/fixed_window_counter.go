package problems

import (
	"fmt"
	"sync"
	"time"
)

type FixedWindow struct {
	capacity       int
	tokens         int
	refillRate     time.Duration
	lastRefillTime time.Time
	mu             sync.Mutex
}

func NewFixedWindow(capacity int, refillRate time.Duration) *FixedWindow {
	return &FixedWindow{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
	}
}

func (fw *FixedWindow) Allow() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	now := time.Now()

	if now.After(fw.lastRefillTime.Add(fw.refillRate)) {
		fmt.Println("here")

		fw.tokens = fw.capacity
		fw.lastRefillTime = now
	}

	if fw.tokens == 0 {
		return false
	}

	fw.tokens--

	return true
}

func FixedWindowEx() {
	fw := NewFixedWindow(5, 2*time.Second)

	for range 10 {
		if fw.Allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}

		time.Sleep(399 * time.Millisecond)
	}
}
