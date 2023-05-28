package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	// множество - это мапа, но только с 2 значениями(есть/нету key в множестве)
	set := make(map[string]struct{})
	for _, v := range data {
		set[v] = struct{}{}
	}
	fmt.Println(set)
}
