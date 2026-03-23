package main

import "fmt"

//create a channel
//ch := make( chan int)

//send

//ch <- 10

//recivce

//val := <-ch

// func main() {

// 	ch := make(chan int)
// 	go func() {
// 		ch <- 10
// 	}()

// 	val := <-ch

// 	fmt.Println((val))

// }
func main() {

	// message := make(chan string)

	// go func() {

	// 	message <- "ping"
	// }()

	// msg := <-message

	// fmt.Println(msg)

	message := make(chan string, 2)

	message <- "buffred"
	message <- "channel"

	fmt.Println(message)
	fmt.Println(message)
}
