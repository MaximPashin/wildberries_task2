package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	val *atomic.Uint64
}

func (c Counter) Inc() {
	c.val.Add(1)
}

func (c Counter) Get() uint64 {
	return c.val.Load()
}

func main() {
	c := Counter{&atomic.Uint64{}}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.Get())
}
