package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func main() {
	// defer is executed in LIFO order
	defer fmt.Println("World")
	fmt.Println("Hello ")
	myDefer()

	a := 199
	defer displayValue(a)
	a = 100
	fmt.Println("Value of a is: ", a)

	methodDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func displayValue(val int) {
	fmt.Println("\nValue is: ", val)
}

func (u User) defMethod() {
	fmt.Printf("Name: %v\nEmail: %v\n", u.Name, u.Email)
}

func methodDefer() {
	user := User{
		Name:  "Steve Rogers",
		Email: "steverogers@gmail.com",
	}

	defer user.defMethod()
	fmt.Println("Welcome")
}
