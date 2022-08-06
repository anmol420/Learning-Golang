package main

import "fmt"

const LoginToken string = "jjsksajkjl" //public or global variables

func main() {
	var username string = "Anmol"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLogged bool = false
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type: %T \n", isLogged)

	var small int = 255
	fmt.Println(small)
	fmt.Printf("Variable is of type: %T \n", small)

	var smallFloat float64 = 25.448466464894465
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and aliases
	var anothervariable string
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type: %T \n", anothervariable)

	// implicit way of declaring variable
	var web = "google.com"
	fmt.Println(web)

	// no var style
	numberOfUser := 300
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
