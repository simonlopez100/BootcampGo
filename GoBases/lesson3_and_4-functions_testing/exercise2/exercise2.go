package main

import "fmt"

func main() {
	fmt.Println("Hello, teacher!")
	var studentFinalNote float64 = calculateAverageNote(7, 8, 9, -3, 10)
	fmt.Printf("Student final note: %.2f\n", studentFinalNote)
}

func calculateAverageNote(grades ...float64) (averageNote float64) {

	gradesSumatory := 0.0
	for _, note := range grades {

		if note < 0 {
			fmt.Println("Note can't be negative")
			return 0.0
		} else {
			gradesSumatory += note
		}
	}
	averageNote = gradesSumatory / float64(len(grades))
	return averageNote
}
