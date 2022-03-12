package conditional

import "fmt"

func Score(result int) {
	score := result
	var grade string

	if score > 75 {
		grade = "Lulus"
	} else if score == 75 {
		grade = "Rata-rata"
	} else if score < 75 {
		grade = "Tidak Lulus"
	} else {
		grade = "Tidak Valid"
	}

	fmt.Println("Grade:", grade)
}