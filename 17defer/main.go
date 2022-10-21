package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("LIFO") //DEFER WORKS ON LIFO - LAST IN FIRST OUT
	fmt.Println("Defer")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
