package main

import "fmt"

func main() {
	var score int
	fmt.Println("Enter your score: ")
	fmt.Scan(&score)

	switch {
	case score > 100 || score < 0:
		fmt.Println("Invalid score.")
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
}
