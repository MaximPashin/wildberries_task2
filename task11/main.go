package main

import "fmt"

func main() {
	// подсчитываю количество встреч числа в множествах,
	// если 2 - число в пересечении обоих множеств
	data1 := []int{4, -2, 2, 9, 0, 5}
	data2 := []int{10, 3, 5, 8, 0, -2}
	temp := make(map[int]int)
	for _, v := range data1 {
		temp[v] += 1
	}
	for _, v := range data2 {
		temp[v] += 1
	}
	result := make([]int, 0)
	for i, v := range temp {
		if v == 2 {
			result = append(result, i)
		}
	}
	fmt.Println(result)
}
