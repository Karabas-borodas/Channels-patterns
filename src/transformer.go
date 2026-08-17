package main

import (
	"fmt"
	// "sync"
	// "time"
)

func Transform[T any](ch1 <-chan T, cation func(T) T) <-chan T {
	outChanels := make(chan T)

	go func() {
		defer close(outChanels)
		for v := range ch1 {
			outChanels <- cation(v)
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

	mul := func(val int) int {
		return val * val
	}
	//декоратор для канала
	for v := range Transform(ch1, mul) {
		fmt.Println(v)
	}
}
