package main

import "fmt"

func main() {

	num := []int{10, 34, 56}

	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))

	num = append(num, 39, 49)

}
