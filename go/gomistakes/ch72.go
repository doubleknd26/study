package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type Donation struct {
		cond    *sync.Cond
		balance int
	}

	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait()
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}

	go f(10)
	go f(15)

	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}
}
