package main

import (
	"fmt"
	// "sync"
	// "time"
)

func Filter[T any](ch1 <-chan T, cation func(T) bool) <-chan T {
	outChanels := make(chan T)

	go func() {
		defer close(outChanels)
		for v := range ch1 {
			if cation(v) {
				outChanels <- v
			}
		}
	}()

	return outChanels
}

func main() {
	ch1 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	predict := func(val int) bool {
		if val%2 == 0 {
			return true
		} else {
			return false
		}
	}
	//декоратор для канала
	for v := range Filter(ch1, predict) {
		fmt.Println(v)
	}
}
