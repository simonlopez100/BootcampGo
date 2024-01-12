package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("Hello, teacher!\n")

	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

	operationMin, err := operation(minimum)
	operationMax, err := operation(maximum)
	operationAvg, err := operation(average)

	if err != nil {
		fmt.Println(err)
		return
	}

	minGradeValue, err := operationMin(7, 8, 9, -3, 10)
	maxGradeValue, err := operationMax(7, 8, 9, -3, 10)
	averageGradeValue, err := operationAvg(7, 8, 9, 10)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Minimum grade: %.2f\n", minGradeValue)
	fmt.Printf("Maximum grade: %.2f\n", maxGradeValue)
	fmt.Printf("Average grade: %.2f\n", averageGradeValue)
}

func operation(operationType string) (func(grades ...int) (float64, error), error) {
	switch operationType {
	case "minimum":
		return minimumGrade, nil
	case "average":
		return averageGrade, nil
	case "maximum":
		return maximumGrade, nil
	default:
		return nil, fmt.Errorf("operation type %s is not valid", operationType)
	}
}

func minimumGrade(grades ...int) (float64, error) {
	if len(grades) == 0 {
		return 0.0, errors.New("No grades provided")
	}

	minValue := grades[0]
	for _, grade := range grades[1:] {
		if grade < minValue {
			minValue = grade
		}
	}

	return float64(minValue), nil
}

func maximumGrade(grades ...int) (float64, error) {
	if len(grades) == 0 {
		return 0.0, errors.New("No grades provided")
	}

	maxValue := grades[0]
	for _, grade := range grades[1:] {
		if grade > maxValue {
			maxValue = grade
		}
	}

	return float64(maxValue), nil
}

func averageGrade(grades ...int) (float64, error) {
	if len(grades) == 0 {
		return 0.0, errors.New("No grades provided")
	}

	gradesSumatory := 0.0
	for _, grade := range grades {
		if grade < 0 {
			return 0.0, fmt.Errorf("grade can't be negative for average")
		} else {
			gradesSumatory += float64(grade)
		}
	}
	averageGrade := gradesSumatory / float64(len(grades))
	return averageGrade, nil
}
