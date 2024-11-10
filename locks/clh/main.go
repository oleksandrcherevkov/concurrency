package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	lock := new()
	counter := &Counter{
		lock: lock,
	}
	createProcesses(counter, 20, 100)
}

type CLHLock struct {
	tail atomic.Pointer[CLHNode]
}

type CLHNode struct {
	locked bool
}

func new() *CLHLock {
	lock := &CLHLock{}
	freeNode := &CLHNode{}
	lock.tail.Store(freeNode)
	return lock
}

func (lock *CLHLock) Lock() *CLHNode {
	node := &CLHNode{locked: true}
	tail := lock.tail.Swap(node)
	for tail.locked {
	}
	return node
}

func (lock *CLHLock) Unlock(node *CLHNode) {
	node.locked = false
}

// the same block as for TASLock
type Counter struct {
	n    int
	lock *CLHLock
}

// method that uses lock to perform a concurrent sensitive operation
func (c *Counter) increment() {
	node := c.lock.Lock()
	temp := c.n
	// sleep for more distinctive gap in counting
	// generally represents any other time consuming operation
	// to see in action - comment the lock and unlock lines
	// desired behavior: read(T1) -> read(T2) -> write(T1) -> write(T2)
	time.Sleep(1 * time.Microsecond)
	c.n = temp + 1
	c.lock.Unlock(node)
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
