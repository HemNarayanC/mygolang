package main

import "fmt"

func main() {
	var username string = "I am learning golang"
	fmt.Println(username) //Println doesn't support formatting like %T
	fmt.Printf("Type of variable: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of variable: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type of variable: %T\n", smallVal)
}
