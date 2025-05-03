package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Implement a concurrent-safe token bucket rate limiter in Go.
The limiter should allow a certain number of requests per second, and bursting up to a max capacity.

⸻

✨ Requirements:
	•	Allow() bool
Returns true if the request is allowed, or false if the rate limit has been exceeded.
	•	You must:
	•	Support a fixed refill rate (e.g., 5 tokens per second)
	•	Support a max burst size (e.g., 10 tokens max)
	•	Ensure the implementation is safe under concurrent calls

⸻

📌 Constraints:
	•	Use Go’s time package (e.g., time.Now(), time.Since()).
	•	You may use sync.Mutex for safety — avoid channels unless necessary.
	•	You don’t need to simulate usage — just the data structure and core method.
*/

type TokenBucket struct {
	l          sync.Mutex // lock
	tokens     int        // number of tokens
	maxBurst   int        // maximum number of tokens
	refillRate int        // tokens per second
	lastRefill time.Time
}

func (tb *TokenBucket) Allow() bool {
	tb.l.Lock()
	defer tb.l.Unlock()

	timeElapsed := time.Since(tb.lastRefill).Seconds()
	tokensToAdd := int(timeElapsed * float64(tb.refillRate))

	if tokensToAdd > 0 {
		tb.tokens = min(tb.maxBurst, tb.tokens+tokensToAdd)
		tb.lastRefill = time.Now()
	}

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("hello, world")
}
