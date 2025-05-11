package main

import (
	"fmt"
	"sort"
)

func main() {
	var pLanguagesList = []string{"C#", "Rust", "Java"}
	fmt.Println(pLanguagesList)

	pLanguagesList = append(pLanguagesList, "JavaScript", "Python", "Go")
	fmt.Println(pLanguagesList)

	// pLanguagesList = append(pLanguagesList[1:])
	// fmt.Println(pLanguagesList)

	// pLanguagesList = append(pLanguagesList[:3])
	// fmt.Println(pLanguagesList)

	pLanguagesList = append(pLanguagesList[1:3])
	fmt.Println(pLanguagesList)

	highScores := make([]int, 4)
	highScores[0] = 456
	highScores[1] = 895
	highScores[2] = 754
	highScores[3] = 645

	fmt.Println(highScores)
	highScores = append(highScores, 485, 987, 254, 684)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	lowScores := make([]int, 4)
	lowScores[0] = 15
	lowScores[1] = 45
	lowScores[2] = 9
	lowScores[3] = 78
	fmt.Println(lowScores)

	lowScores = append(lowScores, 12, 16, 85)
	fmt.Println(lowScores)

	var sliceLowScores = sort.IntSlice(lowScores)
	fmt.Println(sliceLowScores)

	sort.Ints(lowScores)
	fmt.Println(lowScores)

	var bookList = []string{"Networking", "DotNet", "MIS", "IOM", "Computer_Graphics_&_Animation"}
	fmt.Println(bookList)
	var index int = 2
	bookList = append(bookList[:index], bookList[index+1:]...)
	fmt.Println(bookList)

}
