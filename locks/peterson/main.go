package main

import (
	"fmt"
	"math/rand"
	"time"
)

var lock = PetersonLock{}

func main() {
	go useSharedFunction(0, 10)
	go useSharedFunction(1, 10)
	time.Sleep(100 * time.Second)
}
func useSharedFunction(threadId int, n int) {
	for i := 0; i <= n; i++ {
		randomSleep(1, 4)
		sharedFunction(threadId)
	}
}
func sharedFunction(threadId int) {
	lock.lock(threadId)
	fmt.Printf("%v CS\n", threadId)
	lock.unlock(threadId)
}

type Lock interface {
	lock(id int)
	unlock(id int)
}

type PetersonLock struct {
	threads [2]bool
	victim  int
}

func (lock *PetersonLock) lock(id int) {
	lock.threads[id] = true
	lock.victim = id
	for lock.victim == id && lock.threads[1-id] {
	}
}

func (lock *PetersonLock) unlock(id int) {
	lock.threads[id] = false
}

func randomSleep(min, max int) {
	d := max - min
	r := rand.Intn(d)
	time.Sleep(time.Duration(r+min) * time.Second)
}
