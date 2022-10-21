package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome To Functions In Golang")
	greeter()

	result := adder(5, 6)
	fmt.Println("Result Is:", result)

	total := proAdder(2, 5, 7, 6)
	fmt.Println("ProAdder Result:", total)
}

func adder(valOne int, ValTwo int) int {
	return valOne + ValTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Namaste")
}
