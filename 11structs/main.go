package main

import "fmt"

func main() {
	fmt.Println("Structs")

	myinfo := User{Name: "Anmol", Email: "xxx@gmail.com", Status: true, Age: 18}
	fmt.Println(myinfo)
	fmt.Printf("Details: %+v\n", myinfo)
	fmt.Printf("Name is %v.", myinfo.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
