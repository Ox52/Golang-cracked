package main

import "fmt"


func main(){


	for i:=0; i<=10; i++ {

		fmt.Println(i);
	  }
	  
	
	for n:= range 10 {
		
		if n%2 == 0{
			fmt.Println(n, "is even")
		}else{
			fmt.Println(n ,"is odd")
		}
	  }
}