package main

import (
	"fmt"
	"sync"
)

type pair struct {
	i int
	v int
}

func main() {
	// 1-ый вариант: обрабатываю каждое элемент массива в отдельной горутине
	// синхронизация между горутинами не нужна, т.к. у каждая горутина работает
	// на своём участке памяти. Единственное нужно синхронизировать обработку
	// массива и вывод, что происходит через WaitGroup
	data := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for i := range data {
		wg.Add(1)
		go func(i int) {
			data[i] = data[i] * data[i]
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(data)

	// 2-ой вариант: через несколько горутин (в данном случаи 2) получающие данные
	// через канал и возвращающие данные через другой канал
	data = []int{2, 4, 6, 8, 10}
	wg = sync.WaitGroup{}
	dataChan := make(chan pair, 5)
	resChan := make(chan pair, 5)
	// запуск горутин
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			// горутины в цикле до закрытия входного канала обрабатывают данные
			// и отправляют результат в выходной канал
			for {
				d, ok := <-dataChan
				if ok {
					d.v *= d.v
					resChan <- d
				} else {
					break
				}
			}
			wg.Done()
		}()
	}
	// горутина воркер не может закрыть не зная о состояниях других горутин
	// поэтому я создаю отдельную которая дожидается исполнения всех остальных
	// и закрывает выходной канал
	go func() {
		wg.Wait()
		close(resChan)
	}()
	// запись данных во входной канал
	for i, v := range data {
		dataChan <- pair{i, v}
	}
	close(dataChan)
	// запись результата обработки
	for {
		d, ok := <-resChan
		if !ok {
			break
		}
		data[d.i] = d.v
	}
	fmt.Println(data)
}
