package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("We gonna study time on go...")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("15:04:05"))
	fmt.Println(presentTime.Format("Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2014, time.April, 8, 16, 5, 80, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("Monday"))
	fmt.Println(createdDate.Format("01-02-2006"))
	fmt.Println(createdDate.Format("15:04:05"))

}
