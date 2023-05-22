package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	fmt.Println("Enter number of workers:")
	var n uint
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err.Error())
	}
	dataChan := make(chan any)
	wg := sync.WaitGroup{}
	for i := uint(0); i < n; i++ {
		wg.Add(1)
		go func(i uint) {
			for d := range dataChan {
				fmt.Println(i+1, " worker process data: ", d)
			}
			wg.Done()
		}(i)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	data := uint(0)
Loop:
	for {
		select {
		case <-c:
			close(dataChan)
			break Loop
		default:
			dataChan <- data
			data += 1
		}
	}
	wg.Wait()
}
