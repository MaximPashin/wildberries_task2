package main

import "fmt"

func delete(arr []int, i int) []int {
	if len(arr) <= i {
		return append(arr)
	}
	return append(arr[:i], arr[i+1:]...)
}

func main() {
	arr := []int{3, -1, 5, 11, -12, 8}
	fmt.Println(delete(arr, 2))
}
