package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := any(make(chan int))
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		if reflect.ValueOf(v).Kind() == reflect.Chan {
			fmt.Println("chan")
		} else {
			fmt.Println("other type")
		}
	}
}
