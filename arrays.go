package main

import(
	"fmt"
)

func main(){


	var  a[5]int
	fmt.Println("empt", a)

	a[4]=100

	fmt.Println("set",a)
	fmt.Println("get",a[4])
	fmt.Println(len(a))


	b:= [5]int {1,2,3,4,5}
	fmt.Println(b)

	c:= [...]int{10,20,30,40,50,60}
	fmt.Println(c)

	twi:=[2][3]int{
		{1,3,4},
		{5,6,7},
	}
	fmt.Println(twi)

}