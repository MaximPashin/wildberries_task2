package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	words := strings.Split(str, " ")
	lp := 0
	rp := len(words) - 1
	// попарно меняем слова в массиве
	for lp < rp {
		words[lp], words[rp] = words[rp], words[lp]
		lp += 1
		rp -= 1
	}
	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(ReverseWords(str))
}
