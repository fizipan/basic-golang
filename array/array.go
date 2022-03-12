package array

import "fmt"

func TestArray() {
	// Declaration
	var languages [4]string

	languages[0] = "Go"
	languages[1] = "Python"
	languages[2] = "Java"
	languages[3] = "PHP"


	// Explicit inline
	// persons := []string{"Hafizh", "Maulana", "Yusliansyah"}

	// Explicit mutiline
	// persons := [...]string{
	// 	"Hafizh", 
	// 	"Maulana", 
	// 	"Yusliansyah",
	// 	"Mantap",
	// }

	// numbers := []int{1, 2, 3, 4, 5}


	fmt.Println("Panjang Array : ", len(languages))

	for i, language := range languages {
		fmt.Println(i + 1, "-", language)
	}
}