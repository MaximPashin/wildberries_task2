package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type pair struct {
	i int
	v int
}

func main() {
	// 1-ый вариант: воркеры работают паралельно каждый на своем участке памяти
	// после завершения всех воркеров в одном потоке агрегируем результат
	data := []int64{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for i := range data {
		wg.Add(1)
		go func(i int) {
			data[i] = data[i] * data[i]
			wg.Done()
		}(i)
	}
	wg.Wait()
	var res int64 = 0
	for _, v := range data {
		res += v
	}
	fmt.Println(res)

	// 2-ый вариант: тк операция сложения ассоциативна и коммутативна мы можем
	// выполнять ее в произвольном порядке. Единственное, нужно что бы сложение
	// выполнялось атомарно.
	data = []int64{2, 4, 6, 8, 10}
	res = 0
	wg = sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v int64) {
			atomic.AddInt64(&res, v*v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println(res)

}
