package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func main() {
	a := Point{1, 1}
	b := Point{5, 4}
	fmt.Printf("Distance between a, b: %v\n", math.Sqrt(math.Pow(float64(a.x-b.x), 2.0)+math.Pow(float64(a.y-b.y), 2.0)))
}
