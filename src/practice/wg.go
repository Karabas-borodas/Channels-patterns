package main

import (
	"fmt"
	"sync"
	_ "time"
)

func workpool(ch1 <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Обработка данных из канала
	for val := range ch1 {
		fmt.Print("ch1-", val, "\n")
		result <- val * 2
	}
}

func main() {
	ch1 := make(chan int, 6)
	result := make(chan int, 6)
	var wg sync.WaitGroup

	// Запускаем 5 воркеров
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workpool(ch1, result, &wg)
	}

	// Отправляем данные в канал
	for i := 5; i < 18; i++ {
		ch1 <- i
	}
	close(ch1)

	// Ждем завершения всех воркеров
	go func() {
		wg.Wait()
		close(result)
	}()

	// Читаем результаты
	for val := range result {
		fmt.Println("Done -", val)
	}
}
