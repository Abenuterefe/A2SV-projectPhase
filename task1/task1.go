package main

import (
	"fmt"
)

func main() {
	studentName := getStudentName()
	subjectCount := getSubjectCount()
	grades := getGrades(subjectCount)
	average := calculateAverage(grades)

	displayResults(studentName, grades, average)
}

func getStudentName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func getSubjectCount() int {
	var count int
	for {
		fmt.Print("Enter the number of subjects: ")
		_, err := fmt.Scanln(&count)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if count > 0 {
			return count
		}

		fmt.Println("Please enter a number greater than zero.")
	}
}

func getGrades(count int) map[string]float64 {
	grades := make(map[string]float64)

	for i := 0; i < count; i++ {
		var subject string
		var grade float64

		fmt.Printf("Enter name of subject #%d: ", i+1)
		fmt.Scanln(&subject)

		for {
			fmt.Printf("Enter grade for %s (0.0 - 100.0): ", subject)
			fmt.Scanln(&grade)
			if grade >= 0 && grade <= 100 {
				break
			}
			fmt.Println("Invalid grade. Please enter a value between 0 and 100.")
		}

		grades[subject] = grade
	}

	return grades
}

func calculateAverage(grades map[string]float64) float64 {
	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func displayResults(name string, grades map[string]float64, average float64) {
	fmt.Printf("\nStudent Name: %s\n", name)
	fmt.Println("Grades:")
	for subject, grade := range grades {
		fmt.Printf("  %s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2f\n", average)
}
