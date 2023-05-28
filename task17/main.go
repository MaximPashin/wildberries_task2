package main

import "fmt"

func BinSearch(arr []int, v int) (int, bool) {
	lp := 0
	rp := len(arr)
	// критерий остановки
	for lp < rp {
		// число найдено
		if arr[(lp+rp)/2] == v {
			return (lp + rp) / 2, true
		}
		// изменение границ поиска в зависимости от того левее число или правее
		if arr[(lp+rp)/2] < v {
			lp = (lp+rp)/2 + 1
		} else {
			rp = (lp+rp)/2 - 1
		}
	}
	if arr[lp] == v {
		return lp, true
	}
	return 0, false
}

func main() {
	arr := []int{-4, -1, 0, 3, 12, 38, 56}
	pos, ok := BinSearch(arr, 12)
	if ok {
		fmt.Println(pos)
	} else {
		fmt.Println("Not found")
	}
}
