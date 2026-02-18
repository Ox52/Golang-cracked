package main


import (

	"fmt"
	
)

func main(){


	var s []string
	fmt.Println("id", s,s==nil,len(s)==0)

	s =make([]string,3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0]="a"
	s[1]="b"
	s[2]="c"


	fmt.Println("set",s)

	s =append(s,"d")
	s =append(s ,"e","f")

	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy:", c)


}