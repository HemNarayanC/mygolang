package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate my code of go")

	//comma ok, err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for rating,", input)
	fmt.Printf("Type of this rating is %T\n", input)
}
