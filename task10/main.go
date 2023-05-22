package main

import (
	"fmt"
	"math"
)

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	hashmap := make(map[int]([]float64))
	for _, v := range data {
		bucket := int(math.Floor(v / 10))
		hashmap[bucket] = append(hashmap[bucket], v)
	}
	for k, v := range hashmap {
		fmt.Println(k*10, ": ", v)
	}
}
