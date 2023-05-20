package main

import "fmt"

// структура Human (с произвольным набором полей и методов)
type Human struct {
	data any
}

func (h Human) SomeMethod() {
	fmt.Println(h.data)
}

// встраивание методов в структуре Action от родительской структуры Human
// встраивание я реализовывал через embedding, таким образом я наследую
// методы 'родительского' класса, не реализуя их
type Action struct {
	Human
	newdata any
}

func (a Action) NewMethod() {
	fmt.Println(a.newdata)
}

func main() {
	a := Action{Human{"Hello"}, "World"}
	a.SomeMethod()
	a.NewMethod()
}
