package main

import "fmt"

func main() {
	// 64 бита - вмещают результаты делений, перемножений и т.д. чисел > 2^20
	var a, b float64
	fmt.Println("Enter two numbers(a,b):")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Printf("a*b = %v\n", a*b)
	fmt.Printf("a/b = %v\n", a/b)
	fmt.Printf("a-b = %v\n", a-b)
	fmt.Printf("a+b = %v\n", a+b)
}
