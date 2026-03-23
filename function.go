package main

import "fmt"

func hello() {

	fmt.Println("hello world")
}

func greet(name string) {

	fmt.Println("hello", name)
}

func add(a int, b int) int {

	return a + b
}

// func add(a, b int) int {
// return a + b
// }
//

func divide(a, b int) (int, int) {

	q := a / b
	r := a % b
	return q, r

}

func main() {

	hello()
	greet("hasrh")

	result := add(4, 5)

	fmt.Println(result)

	q, r := divide(10, 3)

	fmt.Println(q, r)

	// q, _ := divide(9, 3) to ignore the value

}
