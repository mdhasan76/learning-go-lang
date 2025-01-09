package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Pointer class")

	// exNum := time.Now()
	// fmt.Println("ExNum variable", exNum)

	// var ptr *int
	// fmt.Println("Value of pointer is:", ptr)

	myNumber := 39
	var ptr *int = &myNumber
	myNumber = myNumber + 1
	fmt.Println("This is actual pointer is:", ptr)
	fmt.Println("This is actual pointer is:", *ptr)
	
	*ptr = 8 + 2
	fmt.Println("This is later pointer is:", ptr)
	fmt.Println("This is later pointer is:", *ptr)
}