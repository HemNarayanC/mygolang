package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//using range with loop
	for day := range days {
		fmt.Println(days[day])
	}

	for index, day := range days {
		fmt.Printf("For index %v day is %v\n", index+1, day)
	}

	for _, day := range days {
		fmt.Printf("day is %v\n", day)
	}
}
