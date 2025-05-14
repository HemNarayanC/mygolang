package main

import "fmt"

func showFallThrough() {
	level := 2 // 1 = Bronze, 2 = Silver, 3 = Gold

	fmt.Print("Your benefits:\n")

	switch level {
	case 3:
		fmt.Println("- Access to VIP Lounge")
		fallthrough
	case 2:
		fmt.Println("- 20% Discount on Flights")
		fallthrough
	case 1:
		fmt.Println("- Free Checked Baggage")
	default:
		fmt.Println("No membership benefits")
	}
}

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

	showFallThrough()
}
