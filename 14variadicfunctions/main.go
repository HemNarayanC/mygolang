package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Printf("Found %v at index %v\n", num, i)
			found = true
		}
	}

	if !found {
		fmt.Printf("Number %v not found\n", num)
	}
	fmt.Printf("\n")
}

func main() {
	find(63, 54, 69, 85, 63, 12)
	find(1, 92, 3, 24, 5, 1, 56, 88, 9)
	find(76, 99, 65, 25, 76, 99, 65, 25, 76, 99, 65, 25)
	find(20)
}
