package main

import (
	"fmt"
	"sync"
	// "time"
	"strconv"
)

func Out[T any](ch1 chan T) []chan T {
	var count int = 2
	Out1 := make([]chan T, count)
	for i := 0; i < count; i++ {
		Out1[i] = make(chan T)
	}
	// var iter = 0
	go func() {
		for v := range ch1 {
			for chanels := range Out1 {
				Out1[chanels] <- v
			}
		}
		for _, ch := range Out1 {
			close(ch)
		}
	}()
	return Out1
}

func main() {

	chan1 := make(chan string)

	var wg sync.WaitGroup

	go func() {
		defer close(chan1)
		for i := 0; i < 10; i++ {
			chan1 <- strconv.Itoa(i)
		}

	}()

	out := Out(chan1)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for va := range out[0] {
			fmt.Println("ch1", va)
		}
	}()
	go func() {
		defer wg.Done()
		for va := range out[1] {
			fmt.Println("ch2", va)
		}
	}()
	wg.Wait()
}
