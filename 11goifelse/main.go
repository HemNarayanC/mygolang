package main

import "fmt"

func checkEligibility(age int) {
	if age < 0 {
		fmt.Println("Provided age is invalid.")
	} else if age < 18 {
		fmt.Println("You are not eligible to vote.")
	} else if age >= 18 && age < 60 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("Age exceeds valid range.")
	}
}

func oddEven() {
	var number int
	fmt.Println("Enter a number: ")
	fmt.Scan(&number)

	if num := number; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func main() {
	var age int
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	checkEligibility(age)

	oddEven()
}
