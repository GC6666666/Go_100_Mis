package No_74

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	counter := NewCounter()
	go func() {
		counter.Increment("foo")
	}()
	go func() {
		counter.Increment("bar")
	}()
}

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{
		counters: make(map[string]int),
	}
}

func (c Counter) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}
