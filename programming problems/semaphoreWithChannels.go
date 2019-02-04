package main 

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	capacity 	 int
	counter  	 int
	waitersLock  sync.Mutex
	waiters	 	 list.List
}

func (s *Semaphore) Acquire() {
	s.waitersLock.Lock()

	if (s.counter > 0) {
		s.counter--

		s.waitersLock.Unlock()
	} else {
		ch := make(chan bool)
		s.waiters.PushBack(ch)

		s.waitersLock.Unlock()

		<- ch
	}
}

func (s *Semaphore) Release() {
	s.waitersLock.Lock()

	if (s.counter < s.capacity) {
		s.counter++

		el := s.waiters.Front()
		if (el != nil) {

			ch := s.waiters.Remove(el).(chan bool)
			ch <- true	
		}
	}

	s.waitersLock.Unlock()
}

func compute(tid int, s *Semaphore) {
	fmt.Printf("Worker #%d preparing....\n", tid)

	s.Acquire()
	fmt.Printf("Worker #%d running....\n", tid)
	time.Sleep(3 * time.Second)
	fmt.Printf("Worker #%d releasing mutex....\n", tid)
	s.Release()

	fmt.Printf("Worker #%d exiting....\n", tid)
}

func main() {
	fmt.Println("Done!")

	s := &Semaphore{}
	s.capacity = 4
	s.counter  = 1

	for i:=0; i <= 8; i++ {
		go compute(i, s)
	}

	time.Sleep(120 * time.Second)
}
