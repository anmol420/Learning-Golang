package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome To Files in Golang")

	content := "To Be Pushed Into File"

	file, err := os.Create("./test.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println(length)
	defer file.Close()
	readFile("./test.txt")
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println(string(data))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
