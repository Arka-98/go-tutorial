package problems

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   int
	tokens     float32
	lastRefill time.Time
	refillRate time.Duration
	mu         sync.Mutex
}

func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     float32(capacity),
		lastRefill: time.Now(),
		refillRate: refillRate,
	}
}

func (tokenBucket *TokenBucket) Allow() bool {
	tokenBucket.mu.Lock()
	defer tokenBucket.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tokenBucket.lastRefill)

	tokenBucket.tokens += float32(elapsed.Seconds() / tokenBucket.refillRate.Seconds())

	if tokenBucket.tokens > float32(tokenBucket.capacity) {
		tokenBucket.tokens = float32(tokenBucket.capacity)
	}

	tokenBucket.lastRefill = now

	if tokenBucket.tokens >= 1 {
		tokenBucket.tokens--

		return true
	}

	return false
}

func TokenBucketEx() {
	tokenBucketInst := NewTokenBucket(5, 500*time.Millisecond)

	for range 10 {
		if tokenBucketInst.Allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}

		time.Sleep(200 * time.Millisecond)
	}
}
