package main

import "fmt"

func main() {
	fmt.Println("Maps")

	lang := make(map[string]string)

	lang["js"] = "JavaScript"
	lang["py"] = "Python"
	lang["go"] = "Golang"

	fmt.Println(lang)
	fmt.Println(lang["js"])

	delete(lang, "py")
	fmt.Println(lang)

	//loops

	for key, value := range lang {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}
}
