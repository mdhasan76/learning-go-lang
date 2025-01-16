package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("world")
	fmt.Println("Deffer class")
	deferLoop()
}
func deferLoop(){
	for i :=1; i <= 5; i++{
		defer fmt.Println(i)
	}
}