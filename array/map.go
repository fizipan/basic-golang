package array

import "fmt"

func TestMap() {
	// Declaration
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["satu"] = 1
	// myMap["dua"] = 2

	// Explicit
	myMap := map[string]int{
		"satu": 1, 
		"dua": 2,
		"tiga": 3,
	}

	for key, value := range myMap {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	fmt.Println("=====================")

	delete(myMap, "satu")
	
	// for key, value := range myMap {
	// 	fmt.Println("Key: ", key, "Value: ", value)
	// }

	_, isAvailable := myMap["empat"]
	fmt.Println(isAvailable)
}