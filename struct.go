package main

import "fmt"

// type user struct {
// 	name string
// 	age  int
// }

// func main() {

// 	u := user{

// 		name: "harsh",
// 		age:  22,
// 	}

// 	fmt.Println(u)

// }
//
type User struct {
	name string
}

func( u User) change(){

	u.name = "aman"
}


func main(){

	u := User( name:"harsh")
   u.change()
   fmt.Println(u.name)

}
