package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Slices")

	var list = []string{"Apple","Peach"}
	fmt.Println(list)

	list = append(list, "Mango", "Banana")
	fmt.Println(list)

	list = append(list[1:3])
	fmt.Println(list)

	highScore := make([]int, 4)

	highScore[0] = 2
	highScore[1] = 4
	highScore[2] = 8
	highScore[3] = 9
	//highScore[4] = 7
	highScore = append(highScore, 11, 10, 12)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)

	//how to remove a value from slice
	var courses = []string{"react","js","java","python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}