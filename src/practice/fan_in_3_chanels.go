package main

import (
	"fmt"
	"sync"
)

func MergeChannel(ch1 chan int, ch2 chan int, ch3 chan int) chan int {
	var wg sync.WaitGroup
	resultChan := make(chan int)

	readFromChannel := func(ch <-chan int) {
		defer wg.Done()
		for v := range ch {
			resultChan <- v
		}
	}
	wg.Add(3)
	go readFromChannel(ch1)
	go readFromChannel(ch2)
	go readFromChannel(ch3)
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

		for i := 0; i < 100; i++ {
			ch1 <- i
			ch2 <- i + 2
			ch3 <- i + 3
		}
	}()
	for value := range MergeChannel(ch1, ch2, ch3) {
		fmt.Println(value)
	}
}
