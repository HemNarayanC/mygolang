package main

import "fmt"

func main() {
	//defer is executed in LIFO order
	defer fmt.Println("World")
	fmt.Println("Hello ")
	myDefer()

	a := 199
	defer displayValue(a)
	a = 100
	fmt.Println("Value of a is: ", a)

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func displayValue(val int) {
	fmt.Println("\nValue is: ", val)
}
