package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of my mood:")

	// comma ok || err err
	hereIsInput, _ := inputReader.ReadString('\n')
	// fmt.Printf(hereIsInput)
	if(hereIsInput == "hello"){
		fmt.Println("Inner true", hereIsInput)
	}
	fmt.Println("Thanks for rating for my mind: ", hereIsInput)
	fmt.Printf("Thanks for rating for my mind type: %T ", hereIsInput)
	// fmt.Printf("Thanks for rating for my mind: not only",)
}