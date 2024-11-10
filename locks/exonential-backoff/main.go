package main

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/oleksandrcherevkov/concurrency/internal/benchmark"
)

func main() {
	lock := &BackoffLock{}
	benchmark.PrintCount(lock, 20, 100)
}

var (
	MinDelay = 1
	MaxDelay = 1 << 4
)

type BackoffLock struct {
	locked atomic.Bool
}

func (lock *BackoffLock) Lock() {
	delay := MinDelay
	for true {
		for lock.locked.Load() {
		}
		if !lock.locked.Swap(true) {
			return
		}
		time.Sleep(time.Duration(rand.Int()%delay) * time.Microsecond)
		if delay < MaxDelay {
			delay = delay * 2
		}
	}
}

func (lock *BackoffLock) Unlock() {
	lock.locked.Store(false)
}
