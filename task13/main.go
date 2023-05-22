package main

import "fmt"

func main() {
	a := 5
	b := 3
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
