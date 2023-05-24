package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	str = strings.ToLower(str)
	temp := make(map[rune]int)
	res := true
	for _, ch := range str {
		temp[ch] += 1
		if temp[ch] > 1 {
			res = false
		}
	}
	fmt.Println(res)
}
