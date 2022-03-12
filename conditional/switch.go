package conditional

import "fmt"

func PrintNumber(number int) {

	switch number {
	case 1:
		fmt.Println("One")
		break
	case 2:
		fmt.Println("Two")
		break
	case 3:
		fmt.Println("Three")
		break
	default:
		fmt.Println("zero")
		break
	}
}