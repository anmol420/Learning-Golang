package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome To Time Study")

	present := time.Now()
	fmt.Println(present)

	//To Print The Date Only
	fmt.Println(present.Format("01-02-2006"))

	//To Print The Day Only
	fmt.Println(present.Format("Monday"))

	//To Print The Time Only
	fmt.Println(present.Format("15:04:05"))

	//To Create A Date

	createdDate := time.Date(2020, time.December, 10, 22, 22, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}