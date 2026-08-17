package main

import (
	"fmt"
	"sync"
	// "time"
)

func SplitChannel[T any](ch1 <-chan T, count int) []<-chan T {
	outChanels := make([]chan T, count)
	for i := 0; i < count; i++ {
		outChanels[i] = make(chan T)
	}

	go func() {
		for v := range ch1 {
			//NOTE: блокирующая запись в канал
			for i := 0; i < count; i++ {
				outChanels[i] <- v
			}
		}
		for _, ch := range outChanels {
			close(ch)
		}
	}()
	//NOTE:только для однонаправленных каналов
	resultCh := make([]<-chan T, count)
	for i := 0; i < count; i++ {
		resultCh[i] = outChanels[i]
	}

	return resultCh
}

func main() {
	ch1 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	split := SplitChannel(ch1, 2)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for value := range split[0] {
			fmt.Println("ch1: ", value)
		}
	}()

	go func() {
		defer wg.Done()
		for value := range split[1] {
			fmt.Println("ch2: ", value)
		}
	}()
	wg.Wait()

}
