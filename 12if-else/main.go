package main

import "fmt"

func main() {
	fmt.Println("If-Else In Golang")

	loginCount := 25
	var result string

	if loginCount < 15 {
		result = "Regular"
	} else if loginCount > 15 {
		result = "Special"
	} else {
		result = "Medium"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Less")
	} else {
		fmt.Println("More")
	}
}
