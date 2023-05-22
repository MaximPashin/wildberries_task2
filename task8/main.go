package main

import (
	"fmt"
	"log"
)

func main() {
	var number int64
	var i uint
	var flag uint
	fmt.Println("Enter int64 number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Enter bit: ")
	_, err = fmt.Scan(&i)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Enter flag: ")
	_, err = fmt.Scan(&flag)
	if err != nil {
		log.Fatal(err.Error())
	}
	if flag != 0 && flag != 1 {
		log.Fatal("flag not bit")
	}
	fmt.Println(number | int64(flag<<i))
}
