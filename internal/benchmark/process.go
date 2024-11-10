package benchmark

import (
	"github.com/oleksandrcherevkov/concurrency/internal/benchmark/structures"
	"github.com/oleksandrcherevkov/concurrency/internal/locks"
)

func PrintCount(lock locks.Lock, threads int, count int) {
	counter := structures.NewCounter(lock)
	createProcesses(counter, threads, count)
}

func createProcesses(counter *structures.Counter, threads int, count int) {
	finish := make(chan bool, threads)
	for i := 0; i < threads; i++ {
		go counter.Count(count, finish)
	}
	for i := 0; i < threads; i++ {
		<-finish
	}
}
