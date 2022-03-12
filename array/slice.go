package array

import "fmt"

func TestSlice() {
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "Playstation 5")
	gamingConsoles = append(gamingConsoles, "Nitendo Switch")
	gamingConsoles = append(gamingConsoles, "Xbox One")

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}

	// fmt.Println(gamingConsoles)
}