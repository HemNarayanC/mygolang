package main

import "fmt"

func main() {
	var fruitList [3]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Pear"

	fmt.Println(fruitList)
	fmt.Println("Length of fruitList is", len(fruitList))

	var cars = [4]string{"Lamborghini", "Ferrari", "Rolls Royce", "Alfa Romeo"}
	fmt.Println("Car list : ", cars)
	fmt.Println("Length of cars: ", len(cars))
}
