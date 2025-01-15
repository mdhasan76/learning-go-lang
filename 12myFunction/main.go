package main

import "fmt"

func main() {
	fmt.Println("Welcome to function tutorial")

	greeting()
}
func greeting(){
	fmt.Println("Hello inner function")
}