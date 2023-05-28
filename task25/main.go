package main

import (
	"fmt"
	"runtime"
	"time"
)

func sleep(interval time.Duration) {
	// реализация sleep, через spinlock проверяющий время
	before := time.Now().Add(interval)
	for before.After(time.Now()) {
		runtime.Gosched()
	}
}

func main() {
	go func() {
		a := 0
		for {
			a += 1
		}
	}()
	fmt.Printf("Start: %v\n", time.Now())
	sleep(10 * time.Second)
	fmt.Printf("End: %v\n", time.Now())
}
