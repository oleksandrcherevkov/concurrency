package main

import (
	"sync/atomic"

	"github.com/oleksandrcherevkov/concurrency/internal/benchmark"
)

func main() {
	lock := &TTASLock{}
	benchmark.PrintCount(lock, 20, 100)
}

type TTASLock struct {
	locked atomic.Bool
}

func (lock *TTASLock) Lock() {
	for true {
		for lock.locked.Load() {
		}
		if !lock.locked.Swap(true) {
			return
		}
	}
}

func (lock *TTASLock) Unlock() {
	lock.locked.Store(false)
}
