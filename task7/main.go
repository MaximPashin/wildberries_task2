package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	outWg := sync.WaitGroup{}
	outChan := make(chan int, 50)
	// горутина для вывода информации из канала
	outWg.Add(1)
	go func() {
		for v := range outChan {
			if v == -666 {
				fmt.Println("====================")
			} else {
				fmt.Println(v)
			}
		}
		outWg.Done()
	}()
	// 1 защищаем map rwmutex-ом
	cmap := make(map[int]int)
	cmap[1] = -1
	m := sync.RWMutex{}
	for i := 0; i <= 250; i++ {
		wg.Add(1)
		go func(i int) {
			if i%5 == 0 {
				m.Lock()
				if cmap[1] < i {
					cmap[1] = i
				}
				m.Unlock()
			} else {
				m.RLock()
				outChan <- cmap[1]
				m.RUnlock()
			}
			wg.Done()
		}(i)
	}

	// ожидание завершения всех горутин
	wg.Wait()
	// вывод элемента содержащегося в map
	outChan <- cmap[1]
	// вывод разделителя
	outChan <- -666

	// 2 thread safe map. Запись может осущетвляться через Load, если нужна просто атомарная запись и через
	// cas, если нужно удостовериться, что на момент записи лежит ожидаемое значение
	smap := sync.Map{}
	smap.Store(1, -1)
	for i := 0; i <= 250; i++ {
		wg.Add(1)
		go func(i int) {
			if i%5 == 0 {
				v, _ := smap.Load(1)
				if v.(int) < i {
					smap.CompareAndSwap(1, v, i)
				}
			} else {
				v, _ := smap.Load(1)
				outChan <- v.(int)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	outChan <- cmap[1]
	close(outChan)
	outWg.Wait()
}
