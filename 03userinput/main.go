package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome To User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter The Rating For Our Pizza:")

	// comma ok || error ok 

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks For Rating:", input)
	fmt.Printf("Type Of This Rating: %T", input)
	fmt.Println()
}
