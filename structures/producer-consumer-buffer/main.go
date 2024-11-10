package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	n := 100
	queue := new(n)
	finish := make(chan bool, 2)
	go read(queue, n, finish)
	go write(queue, n, finish)
	<-finish
	<-finish
}

type Queue struct {
	buffer []int
	head   int
	tail   int
}

func new(size int) *Queue {
	return &Queue{
		buffer: make([]int, size),
	}
}

func (q *Queue) enqueue(n int) error {
	full := q.head-q.tail == len(q.buffer)
	if full {
		return errors.New("queue is full")
	}
	q.buffer[q.tail%len(q.buffer)] = n
	q.tail++
	return nil
}

func (q *Queue) dequeue() (int, error) {
	empty := q.head == q.tail
	if empty {
		return 0, errors.New("queue is empty")
	}
	value := q.buffer[q.head%len(q.buffer)]
	q.head++
	return value, nil
}

func write(queue *Queue, n int, finish chan bool) {
	for i := 0; i < n; i++ {
		time.Sleep(1 * time.Microsecond)
		err := queue.enqueue(i)
		if err != nil {
			fmt.Println(err)
		}
	}
	finish <- true
}

func read(queue *Queue, n int, finish chan bool) {
	for i := 0; i < n; i++ {
		x, err := queue.dequeue()
		if err != nil {
			//fmt.Println(err)
		} else {
			fmt.Println(x)
		}
	}
	finish <- true
}
