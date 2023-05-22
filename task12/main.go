package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, v := range data {
		set[v] = struct{}{}
	}
	fmt.Println(set)
}
