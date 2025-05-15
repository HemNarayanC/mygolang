package main

import "fmt"

func main() {
	defer fmt.Println("\nWorld")
	fmt.Println("Hello ")
	myDefer()

	a := 199
	defer displayValue(a)

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func displayValue(val int) {
	fmt.Println("\nValue is: ", val)
}
