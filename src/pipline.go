package main

import (
	"fmt"
	// "sync"
	// "time"
)

func generate[T any](value ...T) <-chan T {
	outChanels := make(chan T)

	go func() {
		defer close(outChanels)
		for _, v := range value {
			outChanels <- v
		}
	}()

	return outChanels
}

func Process[T any](ch1 <-chan T, action func(T) T) <-chan T {

	outChanels := make(chan T)

	go func() {
		defer close(outChanels)
		for v := range ch1 {
			outChanels <- action(v)
		}
	}()

	return outChanels

}

func main() {

	// ch1 := make(chan int)

	slice := []int{1, 2, 3, 4, 5}

	mul := func(value int) int {
		return value * 1
	}
	for v := range Process(generate(slice...), mul) {
		fmt.Println(v)

	}
}
