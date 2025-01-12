package main

import "fmt"

func main() {
	// a := 10
	// b := 20

	// // sum := a + b
	// result := add(a, b)
	// result2 := add(50, 47)
	// fmt.Println("The sum is", result)
	// fmt.Println("The sum is", result2)
	// fmt.Println("The sum is", add(50, 30))


	// // with return value
	// sumNum,multiplicationNum, stringOfNothing := addAndMultiplication(a , b)
	// fmt.Println("then 2nd fn", sumNum)
	// fmt.Println("then 2nd fn", multiplicationNum)
	// fmt.Println("then 2nd fn", stringOfNothing)

	fmt.Println("Welcome to GO function")
	var name string
	var num1 int
	var num2 int
	fmt.Println("Enter your name",)
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
	fmt.Println("Enter first number")
	fmt.Scanln(&num1)
	fmt.Println("Enter Second number")
	fmt.Scanln(&num2)

	fmt.Printf("The sum of %s number is %d \n",name ,num1 + num2)

}

func addAndMultiplication (num1 int, num2 int)(int, int, string) {
	sum := num1 + num2

	multiplication := num1 * num2
	// fmt.Println("sum", sum)
	// fmt.Println("multiplication", multiplication)
	return sum, multiplication, "hello"
}

func add(num1 int, num2 int) int{
	sum := num1 + num2
	return sum
}