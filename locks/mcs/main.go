package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	lock := new()
	counter := &Counter{}
	createProcesses(counter, lock, 20, 100)
}

// should be passed as struct
//
// necessary as no thread local value for node field
//
// also possible to change methods signature to pass
// variable inside thread but it burden algorithm
// implementation with some language specific techniques
//
// no third party library for thread local value is used
// to reduce code complexity, also it is not clear how
// efficient a potential library would be
type MCSLock struct {
	// double pointer
	//
	tail *atomic.Pointer[MCSNode]

	node *MCSNode
}

type MCSNode struct {
	locked bool
	next   *MCSNode
}

func new() MCSLock {
	lock := MCSLock{}
	lock.tail = &atomic.Pointer[MCSNode]{}
	return lock
}

func (lock *MCSLock) Lock() {
	lock.node = &MCSNode{locked: true}
	tail := lock.tail.Swap(lock.node)
	if tail != nil {
		tail.next = lock.node
		for lock.node.locked {
		}
	}
}

func (lock *MCSLock) Unlock() {
	if lock.node.next == nil {
		if lock.tail.CompareAndSwap(lock.node, nil) {
			return
		}
		for lock.node.next == nil {
		}
	}
	lock.node.next.locked = false
}

// the same block as for TASLock
type Counter struct {
	n int
}

// method that uses lock to perform a concurrent sensitive operation
func (c *Counter) increment(lock MCSLock) {
	lock.Lock()
	temp := c.n
	// sleep for more distinctive gap in counting
	// generally represents any other time consuming operation
	// to see in action - comment the lock and unlock lines
	// desired behavior: read(T1) -> read(T2) -> write(T1) -> write(T2)
	time.Sleep(1 * time.Microsecond)
	c.n = temp + 1
	lock.Unlock()
}

func createProcesses(counter *Counter, lock MCSLock, threads int, n int) {
	finish := make(chan bool, threads)
	for i := 0; i < threads; i++ {
		go count(counter, lock, n, finish)
	}
	for i := 0; i < threads; i++ {
		<-finish
	}
}

func count(counter *Counter, lock MCSLock, n int, finish chan bool) {
	for i := 0; i < n; i++ {
		counter.increment(lock)
		fmt.Println(counter.n)
	}
	finish <- true
}
