package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	threadsNumber := 20
	lock := newAnderson(threadsNumber)
	counter := &Counter{
		lock: lock,
	}
	createProcesses(counter, 20, 100)
}

type AndersonLock struct {
	next        atomic.Int64
	queue       []bool
	queueLength int64
	// assumed thread local variable: myPosition
}

func newAnderson(threads int) *AndersonLock {
	queue := make([]bool, threads)
	queue[0] = true
	return &AndersonLock{
		queue:       queue,
		queueLength: int64(threads),
	}
}

func (lock *AndersonLock) lock() int64 {
	// no getAndIncrement function, so getting previous value in two operations
	// not risky, because shared value modified atomically
	// local value always changed before check
	current := lock.next.Add(1)
	current = current - 1
	for !lock.queue[current%lock.queueLength] {
	}
	lock.queue[current%lock.queueLength] = false
	return current
}

func (lock *AndersonLock) unlock(position int64) {
	lock.queue[(position+1)%lock.queueLength] = true
}

// the same block as for TASLock
type Counter struct {
	n    int
	lock *AndersonLock
}

// method that uses lock to perform a concurrent sensitive operation
func (c *Counter) increment() {
	currentPosition := c.lock.lock()
	temp := c.n
	// sleep for more distinctive gap in counting
	// generally represents any other time consuming operation
	// to see in action - comment the lock and unlock lines
	// desired behavior: read(T1) -> read(T2) -> write(T1) -> write(T2)
	time.Sleep(1 * time.Microsecond)
	c.n = temp + 1
	c.lock.unlock(currentPosition)
}

func createProcesses(counter *Counter, threads int, n int) {
	finish := make(chan bool, threads)
	for i := 0; i < threads; i++ {
		go count(counter, n, finish)
	}
	for i := 0; i < threads; i++ {
		<-finish
	}
}

func count(counter *Counter, n int, finish chan bool) {
	for i := 0; i < n; i++ {
		counter.increment()
		fmt.Println(counter.n)
	}
	finish <- true
}
