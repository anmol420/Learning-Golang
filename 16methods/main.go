package main

import "fmt"

func main() {
	fmt.Println("Methods")

	myinfo := User{Name: "Anmol", Email: "xxx@gmail.com", Status: true, Age: 18}
	fmt.Println(myinfo)
	fmt.Printf("Details: %+v\n", myinfo)
	fmt.Printf("Name is %v.", myinfo.Name)

	myinfo.GetStatus()
	myinfo.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// methods
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "xxxx@gmail.com"
	fmt.Println("New Email:", u.Email)
}
