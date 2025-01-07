package No_72

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDonation(t *testing.T) {
	donation := &Donation{}
	f := func(goal int) {
		donation.mu.RLock()
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d goal reacher\n", donation.balance)
		donation.mu.RUnlock()
	}
	go f(10)
	go f(15)

	go func() {
		for {
			time.Sleep(10 * time.Second)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()
	select {}
}

type Donation struct {
	mu      sync.RWMutex
	balance int
}
