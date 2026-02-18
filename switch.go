package main


import (
	"fmt"
	"time"
)


func main(){

	 i:= 2 

	 fmt.Println(" the number is " ,i)

	 switch i{
	 case 1:

		fmt.Println("one")

	 case 2:

		fmt.Println("two")

	 case 3:
		fmt.Println("threse")

	
	 }


	 switch time.Now().Weekday(){

	 case time.Sunday , time.Saturday:
		fmt.Println("its a weeknd ")

	 default:
		fmt.Println("its weekday")
	 }

	 t:= time.Now()

	 switch{

	 case t.Hour()<12:
		fmt.Println("its bedore noon")

	 default:
		fmt.Println("its after noon")
	 }
}