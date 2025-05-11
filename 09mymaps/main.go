package main

import "fmt"

func main() {
	/*maps is used to store key-value pairs
	Syntax:
			make(map[type of key]type of value)
	*/

	currencyCode := make(map[string]string)
	fmt.Println(currencyCode)

	//adding items to the map
	currencyCode["USD"] = "US Dollar"
	currencyCode["NPR"] = "Nepalese Rupee"
	currencyCode["EUR"] = "Euro"
	currencyCode["QAR"] = "Qatari Riyal"
	fmt.Println(currencyCode)

	//deleting the element
	delete(currencyCode, "NPR")

	fmt.Println(currencyCode)
	fmt.Println(currencyCode["QAR"])

	//key value pair using range to iterate over all elements in map
	for key, value := range currencyCode {
		fmt.Printf("%v = %v\n", key, value)
	}

	cyCode := "USD"
	if currencyName, ok := currencyCode[cyCode]; ok {
		fmt.Printf("%v\n", currencyName)
	} else {
		fmt.Println("Currency doesn't exist")
	}

	initialization()

}
