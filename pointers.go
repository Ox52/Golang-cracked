package main

import "fmt"

func main() {

	x := 5
	p := &x

	*p = 20

	fmt.Println(x)

}
