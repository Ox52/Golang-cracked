package main

import "fmt"

type Animal interface {
	speak()
}

type cat struct{}

func (c cat) speak() {
	fmt.Println("moew")
}

func main() {

	var a Animal = cat{}
	a.speak()
}
