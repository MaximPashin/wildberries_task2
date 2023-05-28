package main

import "fmt"

func main() {
	a := 5
	b := 3
	fmt.Println(a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
