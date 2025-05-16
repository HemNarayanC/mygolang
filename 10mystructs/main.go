package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	status bool
}

func main() {
	user1 := User{"Hemant", "hemant@gmail.com", 23, true}
	fmt.Println(user1)

	//creating struct specifying field names
	user2 := User{
		Name:   "Sudip",
		Email:  "sudip@gmail.com",
		Age:    25,
		status: false,
	}

	user3 := &User{
		Name:   "Omrina Chaudhary",
		Email:  "omrina200@gmail.com",
		Age:    13,
		status: true,
	}

	fmt.Printf("Details of User1 are: %+v\n", user1)
	//accessing the field
	fmt.Println(user1.Name)
	fmt.Printf("Name: %v\nEmail: %v\nAge: %v\n", user1.Name, user1.Email, user1.Age)

	fmt.Printf("Details of User2 are: %+v\n", user2)

	fmt.Printf("Name of user3 is %v", (*user3).Name)
}
