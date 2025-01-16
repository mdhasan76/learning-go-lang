package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods class")

	hasan := User{"Hasan Mia", "mdhasanmiah8064@gmail.com", true, 23}
	fmt.Printf("User name is %v , the email of %v is %v, he is %v years old. and he is and active user. his status is %v \n", hasan.name, hasan.name, hasan.email, hasan.age, hasan.status )
	fmt.Println(hasan)
	// getStatus(hasan)
	hasan.getStatus()
}

type User struct {
	name string
	email string
	status bool
	age int
}

func (user User) getStatus(){
	fmt.Println("The user status is", user.status)
}