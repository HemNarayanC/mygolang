package main

import (
	"fmt"
	"math/rand"
	"time"
)

func showFallThrough() {
	var level int
	fmt.Println("Enter the level: ")
	fmt.Scan(&level)
	//level := 2 // 1 = Bronze, 2 = Silver, 3 = Gold

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

func ludoGame() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("You rolled a %d\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 step")
	case 2:
		fmt.Println("You can move 2 steps")
	case 3:
		fmt.Println("You can move 3 steps")
		// fallthrough is used to execute the next case even if the current case is matched
		fallthrough
	case 4:
		fmt.Println("You can move 4 steps")
		fallthrough
	case 5:
		fmt.Println("You can move 5 steps")
	case 6:
		fmt.Println("You can move 6 steps")
	default:
		fmt.Println("You can roll again")
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
	ludoGame()
}
