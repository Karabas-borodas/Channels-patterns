package main

import (
	"fmt"
	"sync"
	// "time"
)

func MargeChan[T any](chanels ...chan T) <-chan T {

	var wg sync.WaitGroup
	wg.Add(len(chanels))

	resultChan := make(chan T)

	for _, va := range chanels {
		go func(c <-chan T) {
			defer wg.Done()
			for val := range c {
				// time.Sleep(time.Second)
				resultChan <- val
			}
		}(va)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	return resultChan
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer func() {
			defer close(ch1)
			defer close(ch2)
			defer close(ch3)
		}()

		for i := 0; i < 10; i++ {
			ch1 <- i
			// ch2 <- i
			ch3 <- i
		}
		for i := 0; i < 2; i++ {
			// ch1 <- i
			ch2 <- i
			// ch3 <- i
		}
	}()

	for variable := range MargeChan(ch1, ch2, ch3) {
		fmt.Println(variable)
	}
}
