package main

import (
	"fmt"
	"math"
)

func main() {
	user1 := User{
		Name:   "Hemant",
		Email:  "hemant@gmail.com",
		Age:    23,
		status: true,
	}

	fmt.Println(user1)
	user1.NewMail()
	fmt.Printf("Details of user1 are: %+v\n", user1)

	//calculating area of circle
	circle := Circle{
		Radius: 5,
	}
	circleArea := circle.Area()
	fmt.Printf("Area of circle is: %v\n", circleArea)
}

type User struct {
	Name   string
	Email  string
	Age    int
	status bool
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length  int
	Breadth int
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (u User) NewMail() {
	u.Email = "newmail@gmail.com"
	fmt.Println("New email is: ", u.Email)
}
