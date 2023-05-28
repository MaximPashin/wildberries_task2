package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	// приведение символов к одному регистру
	str = strings.ToLower(str)
	temp := make(map[rune]int)
	res := true
	// проверка символов на уникальность, через мапу количества встреч символа
	// если символ встречается более 1 раза, он не уникален
	for _, ch := range str {
		temp[ch] += 1
		if temp[ch] > 1 {
			res = false
		}
	}
	fmt.Println(res)
}
