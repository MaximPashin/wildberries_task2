package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	// 1 - горутина выполняет всю работу и завершается
	wg.Add(1)
	go func() {
		fmt.Println("goroutine work")
		fmt.Println("finishing goroutine")
		wg.Done()
	}()
	wg.Wait()

	// 2 - сигнализируем по каналу об необходимости закончить выполнение горутины
	stop := make(chan any)
	wg.Add(1)
	go func() {
		fmt.Println("goroutine work")
	Loop:
		for {
			select {
			case <-stop:
				break Loop
			default:
				time.Sleep(time.Second)
			}
		}
		fmt.Println("finishing goroutine")
		wg.Done()
	}()
	close(stop)
	wg.Wait()
}
