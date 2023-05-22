package main

import (
	"fmt"
)

func main() {
	data := []int{2, 4, 6, 8, 10}
	dataChan := make(chan int)
	resChan := make(chan int)
	go func() {
		// горутина в цикле до закрытия входного канала обрабатывают данные
		// и отправляют результат в выходной канал
		for d := range dataChan {
			resChan <- d * d
		}
		close(resChan)
	}()
	// горутина записывающая данные во входной канал
	go func() {
		for _, v := range data {
			dataChan <- v
		}
		close(dataChan)
	}()
	// вывод результата обработки
	for v := range resChan {
		fmt.Println(v)
	}
}
