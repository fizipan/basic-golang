package conditional

import "fmt"

func ForLooping() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i <= 10 {
	// 	fmt.Println("Hello Hafizh maulana", i)
	// 	i++
	// }

	// var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// for _, value := range numbers {
	// 	fmt.Println(value)
	// }

	characters := "Hafizh Maulana"

	for _, value := range characters {
		if value % 2 == 0 {
			displayString := string(value)
			fmt.Println(displayString)
		}
	}
}