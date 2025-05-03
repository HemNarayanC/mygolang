package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to conversion:")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any number: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered ", input)

	inputNum, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Hello this is error here")
	} else {
		fmt.Println("After adding input to 7 =", inputNum+7)
	}
}
