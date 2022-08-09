package main

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
	//"time"
)

func main()  {
	fmt.Println("Welcome To Maths")

	//random number

	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5)+1)

	// random from crypto

	myRand, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println("Random Number", myRand)
}