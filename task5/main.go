package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Enter work time(seconds):")
	var n uint
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err.Error())
	}
	dataChan := make(chan any)
	go func() {
		for d := range dataChan {
			fmt.Println(d)
		}
	}()
	data := 0
	timer := time.After(time.Duration(n) * time.Second)
Loop:
	for {
		select {
		case <-timer:
			close(dataChan)
			break Loop
		default:
			dataChan <- data
			data += 1
		}
	}
}
