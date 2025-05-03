package main

import "fmt"

const LoginId string = "HelloGoLang54"

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

	var smallFloat float64 = 329.5454874316654545
	fmt.Println(smallFloat)
	fmt.Printf("Type of variable: %T\n", smallFloat)

	//default values and some alliases
	//if declared but not assigned value then by default 0 is assigned
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Type of variable: %T\n", anotherVariable)

	var str string
	fmt.Println(str)
	fmt.Printf("Type of variable: %T\n", str)

	//implicit type
	//lexer here defines the variable type by itself
	var website = "www.google.com"
	fmt.Println(website)

	//no var style
	numberOfUser := 150 //walrun operator for declaration + assignment and is valid inside method only
	fmt.Println(numberOfUser)

	fmt.Printf("Type of variable is %T\n", LoginId)
	fmt.Println(LoginId)

}
