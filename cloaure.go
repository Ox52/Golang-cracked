package main

import "fmt"

func counter() func() int {

	count := 0

	return func() int {
		count++
		return count
	}
}

// recuriosn

func f(n int) {

	if n == 0 {
		return
	}
	f(n - 1)
	fmt.Println(n)

}

func main() {

	c := counter()

	fmt.Println(c)

	f(4)

}
