package main

import "fmt"

func main()  {
	fmt.Println("Array")

	var list [4]string

	list[0]= "Apple"
	list[1]= "Grape"
	list[2]= "Orange"
	list[3]= "Banana"

	fmt.Println(list)
	fmt.Println(len(list))

	var veglist = [3]string{"Potato","Mushroom","Beans"}

	fmt.Println(veglist)
	fmt.Println(len(veglist))
}