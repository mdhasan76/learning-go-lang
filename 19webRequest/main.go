package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/199"

func main() {
	fmt.Println("Welcome to go handle web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the response type is %T \n", response)
	fmt.Println(response)
	defer response.Body.Close()

	dataByte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataByte)
	fmt.Println(content)
}
