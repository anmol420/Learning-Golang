package main

import "fmt"

func main() {
	fmt.Println("Loop In Golang")

	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	//for loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	//for loop type-2
	for i := range days {
		fmt.Println(days[i])
	}

	//for loop type-3
	for j, day := range days {
		fmt.Printf("Index is %v and the valud is %v\n", j, day)
	}

	rv := 1
	for rv < 10 {

		if rv == 2 {
			goto x //print label x
		}

		if rv == 5 {
			break
		}
		fmt.Println(rv)
		rv++
	}

	//label
x:
	fmt.Println("Processing")
}
