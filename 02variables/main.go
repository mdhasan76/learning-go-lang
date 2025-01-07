package main

import "fmt"
const hello = "hello bro"
//  number2 := uint16(25545)
func main() {
	var userName string = "Hasan"
	const a = "hello"
	fmt.Println(userName, a)
	fmt.Printf("variable type is %T \n", userName  )

	var isBlocked bool = false;
	fmt.Println(isBlocked);
	const isBlocked2 bool = true
	fmt.Println(isBlocked2)
	fmt.Printf("This is type 2 %T \n", isBlocked)
	fmt.Printf("This is type 1 %T \n", isBlocked2)

	
	// var number uint16 =33222 
	// fmt.Println(number)f

	var unAssignedVar int8
	fmt.Println("THis is unassigned",unAssignedVar)

	//  hello := uint16(25545)
	// fmt.Println("THis is hello variable",number2)

	var decideItSelf  = false
	fmt.Printf("this is self assigned type %T", decideItSelf) 
	fmt.Printf("this is self assigned type %T", hello)
}