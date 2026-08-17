package main

import (
	"fmt"
	// "structs"
	// "sync"
	"time"
)

func process(ch1 <-chan struct{}) <-chan struct{} {
	outChanels := make(chan struct{})

	go func() {
		defer close(outChanels)
		for {
			select {
			case <-ch1:
				return
			default:
				fmt.Println("процесс идет") //процессинг
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	return outChanels
}

func main() {
	ch1 := make(chan struct{})
	//декоратор для канала
	closeDne := process(ch1)
	close(ch1)
	<-closeDne
	fmt.Println("terminated")
}
