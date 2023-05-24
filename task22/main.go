package main

import "fmt"

func main() {
	var a, b float64
	fmt.Println("Enter two numbers(a,b):")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Printf("a*b = %v\n", a*b)
	fmt.Printf("a/b = %v\n", a/b)
	fmt.Printf("a-b = %v\n", a-b)
	fmt.Printf("a+b = %v\n", a+b)
}
