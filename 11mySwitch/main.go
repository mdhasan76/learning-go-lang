package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to go switch")
	// rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("This is 1")
	case 2:
		fmt.Println("This is 2")
	case 3:
		fmt.Println("This is 3")
	default:
		fmt.Println("This is default")
	}

	num2 := 3
	// for num2 < 10{
	// 	fmt.Printf("the index is %v\n", num2)
	// 	num2++
	// }

	for i := 1; i < num2;i++{
		fmt.Printf("the index of the loop is %v\n", i)
	}
}