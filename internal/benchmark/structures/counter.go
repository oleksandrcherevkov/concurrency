package structures

import (
	"fmt"
	"time"

	"github.com/oleksandrcherevkov/concurrency/internal/locks"
)

// the same block as for TASLock
type Counter struct {
	n    int
	lock locks.Lock
}

func NewCounter(lock locks.Lock) *Counter {
	return &Counter{
		lock: lock,
	}
}

// method that uses lock to perform a concurrent sensitive operation
func (c *Counter) Increment() {
	c.lock.Lock()
	temp := c.n
	// sleep for more distinctive gap in counting
	// generally represents any other time consuming operation
	// to see in action - comment the lock and unlock lines
	// desired behavior: read(T1) -> read(T2) -> write(T1) -> write(T2)
	time.Sleep(1 * time.Microsecond)
	c.n = temp + 1
	c.lock.Unlock()
}

func (counter *Counter) Count(n int, finish chan bool) {
	for i := 0; i < n; i++ {
		counter.Increment()
		fmt.Println(counter.n)
	}
	finish <- true
}
