package main

import (
	"fmt"
	"sync"
)

func MergeVar[T any](chanels ...<-chan T) <-chan T {
	resultChan := make(chan T)

	var wg sync.WaitGroup
	wg.Add(len(chanels))

	for _, chanels := range chanels {
		go func(ch <-chan T) {
			defer wg.Done()
			for val := range ch {
				resultChan <- val
			}
		}(chanels)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	return resultChan
}

func main() {
	ch1 := make(chan rune)
	ch2 := make(chan rune)
	ch3 := make(chan rune)

	go func() {
		defer func() {
			defer close(ch1)
			defer close(ch2)
			defer close(ch3)
		}()

		for i := 0; i < 10; i++ {
			ch1 <- 'A'
			ch2 <- 'B'
			ch3 <- 'C'
		}

	}()
	for value := range MergeVar(ch1, ch2, ch3) {
		fmt.Println(string(value))
	}

}
