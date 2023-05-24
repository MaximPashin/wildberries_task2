package main

import "fmt"

func Reverse(str string) string {
	target := []rune(str)
	lp := 0
	rp := len(target) - 1
	for lp < rp {
		target[lp], target[rp] = target[rp], target[lp]
		lp += 1
		rp -= 1
	}
	return string(target)
}

func main() {
	str := "главрыба"
	fmt.Println(Reverse(str))
}
