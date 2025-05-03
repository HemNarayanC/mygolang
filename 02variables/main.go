package main

import "fmt"

func main() {
	var username string = "I am learning golang"
	fmt.Println(username) //Println doesn't support formatting like %T
	fmt.Printf("Type of variable: %T", username)
}
