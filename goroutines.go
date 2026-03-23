package main

import (
	"fmt"
	"time"
)

func task(name string) {

	for i := 0; i < 3; i++ {

		fmt.Println(name, i)
	}
}

func main() {

	task("a")
	go task("b")

	time.Sleep(time.Second)

}
