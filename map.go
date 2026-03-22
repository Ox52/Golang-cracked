package main

import "fmt"

func main() {

	// m := map[string]int{

	// 	"hasrh": 29,
	// 	"dad":   45,
	// }

	// ma := make(map[string]int)

	// ma["a"] = 12
	// ma["b"] = 23

	// ma["a"] = 100
	// ma["c"] = 300

	// delete(m, "a")

	// fmt.Println(m["a"])

	// fmt.Println(m)
	//

	//create map
	scores := make(map[string]int)

	//add values
	scores["math"] = 100
	scores["science"] = 95

	fmt.Println("map:", scores["math"])
	fmt.Println("map:", scores["science"])
	//check key
	val, ok := scores["english"]

	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("not found")
	}

	//update value
	scores["math"] = 90
	scores["english"] = 80

	fmt.Println(scores["english"])

	delete(scores, "scince")

	//val,ok

	// scoress := map[string]int{
	// 	"math": 90,
	// }
	// val, ok := scoress["math"]
	// fmt.Println(val, ok)

	// val, ok = scoress["enlisgh"]
	// fmt.Println(val, ok)

	// val, ok := scoressx["english"]

	// if ok {

	// 	fmt.Println("val")
	// } else {

	// 	fmt.Println("nope")
	// }

}
