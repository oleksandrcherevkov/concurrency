package main

import (
	"sync/atomic"

	"github.com/oleksandrcherevkov/concurrency/internal/benchmark"
)

func main() {
	lock := &TASLock{}
	benchmark.PrintCount(lock, 20, 100)
}

type TASLock struct {
	locked atomic.Bool
}

func (lock *TASLock) Lock() {
	for lock.locked.Swap(true) {
	}
}

func (lock *TASLock) Unlock() {
	lock.locked.Store(false)
}
