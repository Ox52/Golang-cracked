package main

import "fmt"

func printvalue[T any](val T) {

	fmt.Println(val)

}

func add[T int | float64](a T, b T) T {

	return a + b
}

func main() {

	printvalue(10)
	printvalue("hasrh")

	fmt.Println(add(2, 4))

}
