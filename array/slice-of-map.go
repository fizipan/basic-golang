package array

import "fmt"

func SliceOfMap() {
	students := []map[string]string{
		{
			"name": "Hafizh",
			"score": "A",
		},
		{
			"name": "Athar",
			"score": "B",
		},
	}

	for _, student := range students {
		fmt.Println(student["name"], "-", student["score"])
	}
}