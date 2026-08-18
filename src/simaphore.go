package main

import (
	"fmt"
	"sync"
	// "sync"
	"time"
)

type Semaphore struct {
	ticket chan struct{}
}

func NewSemaphore(ticketsNumber int) Semaphore {
	return Semaphore{
		ticket: make(chan struct{}, ticketsNumber),
	}
}

func (s *Semaphore) Acquire() {
	s.ticket <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ticket
}

func main() {
	// ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(7)

	semaphore := NewSemaphore(4)
	for i := 0; i < 7; i++ {
		semaphore.Acquire()
		go func() {
			defer func() {
				wg.Done()
				semaphore.Release()
			}()
			fmt.Println("working")
			time.Sleep(1 * time.Second)
			fmt.Println("exist..")
		}()

	}

	wg.Wait()
}
