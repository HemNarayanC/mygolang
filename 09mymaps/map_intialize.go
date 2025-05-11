package main

import "fmt"

func initialization() {
	// fmt.Println("Hello")
	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}
	currencyCode["INR"] = "Indian Rupee"
	fmt.Println(currencyCode)
}
