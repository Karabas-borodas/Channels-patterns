package main

import (
	"fmt"
	"sync"
)

func fun[T any](chanels ...<-chan T) chan T {
	resultChan := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(chanels))
	for _, v := range chanels {
		go func() {
			defer wg.Done()
			for va := range v {
				resultChan <- va
			}
		}()

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
			ch2 <- i * 2
			ch3 <- i * 3
		}

	}()
	for variable := range fun(ch1, ch2, ch3) {
		fmt.Println(variable)
	}
}
