package main

import "fmt"

type status int

const (
	pending status = iota
	approved
	rejected
)

func main() {

	var s status = approved

	fmt.Println(s)

}
