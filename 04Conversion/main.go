package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Conversion")

	fmt.Println("Giving feedback for us!")
	 reader := bufio.NewReader(os.Stdin)
	 
	 input, _ := reader.ReadString('\n')
	 fmt.Println("Thanks for giving feedback", input)
	 
	 var takeNum = strings.TrimSpace(input)

	 numRating,err := strconv.ParseFloat(takeNum, 64)
	 fmt.Println(input, numRating, err, "all of the component")
	 if err != nil{
		 fmt.Println(err)
       }  else {
			 fmt.Println("Added 1 to the input", 1 + numRating)
		}

	 
}