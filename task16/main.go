package main

import "fmt"

func QuickSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	// переносим элементы, что бы часть слева была меньше или равна выбранному элементу
	// а правая была строго больше выбранного элемента
	lp := 1
	rp := len(arr) - 1
	for lp < rp {
		if arr[0] >= arr[lp] {
			lp += 1
		}
		if arr[0] < arr[rp] {
			rp -= 1
		}
		if arr[0] < arr[lp] && arr[0] >= arr[rp] && lp < rp {
			arr[lp], arr[rp] = arr[rp], arr[lp]
			lp += 1
			rp -= 1
		}
	}
	if arr[0] >= arr[rp] {
		arr[0], arr[rp] = arr[rp], arr[0]
		rp += 1
	} else {
		arr[0], arr[rp-1] = arr[rp-1], arr[0]
	}
	// проверка на достижение конца рекурсии
	if len(arr) < 3 {
		return
	}
	// рекурсивный вызов на полученных частях
	QuickSort(arr[:rp-1])
	QuickSort(arr[rp:])
}
func main() {
	arr := []int{5, 2, 8, -11, 21, 0, 12, -5}
	QuickSort(arr)
	fmt.Println(arr)
}
