package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	status bool
}

func main() {
	User1 := User{"Hemant", "hemant@gmail.com", 23, true}
	fmt.Println(User1)

	User2 := User{
		Name:   "Sudip",
		Email:  "sudip@gmail.com",
		Age:    25,
		status: false,
	}

	fmt.Printf("Details of User1 are: %+v\n", User1)
	//accessing the field
	fmt.Println(User1.Name)
	fmt.Printf("Name: %v\nEmail: %v\nAge: %v\n", User1.Name, User1.Email, User1.Age)

	fmt.Printf("Details of User2 are: %+v\n", User2)
}
