package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go...")

	x := 10
	// var p *int = &x	//declaration and initialization
	var p = &x

	fmt.Println("Address of x: ", p)
	fmt.Println("Value stored at p: ", *p) //dereferencing

	*p = 15 //changes the value at memory level
	fmt.Println("New value of x: ", x)

}
