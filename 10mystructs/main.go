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
}
