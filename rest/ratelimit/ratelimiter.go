package ratelimit

import (
	"sync"
	"time"
)

// Big thanks to https://github.com/spencersharkey for sharing his ratelimiter with me

type Ratelimiter struct {
	sync.Mutex
	Store RateLimitStore
}

func NewConcurrencyLimiter(store RateLimitStore) *Ratelimiter {
	return &Ratelimiter{
		Store: store,
	}
}

func (l *Ratelimiter) ExecuteCall(bucket string, ch chan error) {
	ttl, err := l.Store.getTTLAndDecrease(bucket)
	if err != nil { // if an error occurred, we should cancel the request
		ch <- err
		return
	}

	if ttl > 0 {
		<-time.After(ttl)
		l.ExecuteCall(bucket, ch)
	} else {
		ch <- nil
	}
}