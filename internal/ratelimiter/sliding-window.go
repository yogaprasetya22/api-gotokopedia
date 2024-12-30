package ratelimiter

import (
	"sync"
	"time"
)

type SlidingWindowRateLimiter struct {
	sync.RWMutex
	clients map[string][]time.Time
	limit   int
	window  time.Duration
}

func NewSlidingWindowLimiter(limit int, window time.Duration) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		clients: make(map[string][]time.Time),
		limit:   limit,
		window:  window,
	}
}

func (rl *SlidingWindowRateLimiter) Allow(ip string) (bool, time.Duration) {
	rl.Lock()
	defer rl.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	// Filter valid timestamps in place (avoid creating a new slice)
	i := 0
	for _, t := range rl.clients[ip] {
		if t.After(windowStart) {
			rl.clients[ip][i] = t
			i++
		}
	}
	rl.clients[ip] = rl.clients[ip][:i]

	// Allow request if within limit
	if len(rl.clients[ip]) < rl.limit {
		rl.clients[ip] = append(rl.clients[ip], now)
		return true, 0
	}

	// Calculate retry time
	retryAfter := rl.clients[ip][0].Add(rl.window).Sub(now)
	return false, retryAfter
}

// Cleanup removes stale clients to free memory
func (rl *SlidingWindowRateLimiter) Cleanup() {
	rl.Lock()
	defer rl.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	for ip, timestamps := range rl.clients {
		i := 0
		for _, t := range timestamps {
			if t.After(windowStart) {
				timestamps[i] = t
				i++
			}
		}
		if i == 0 {
			delete(rl.clients, ip)
		} else {
			rl.clients[ip] = timestamps[:i]
		}
	}
}
