package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files class")

	content := "Hello! i am creating txt file using go lang."
	res, err := os.Create("./content.txt")
	checkNilErr(err)
	length, err := io.WriteString(res, content)
	checkNilErr(err)
	fmt.Println(length)
	 res.Close()
	fmt.Println(res)
	readFile("./content.txt")
	// err := os.Remove("content.txt")
	checkNilErr(err)
	filePath := "content.txt"
	
	errOnDeleting := os.Remove(filePath)
	checkNilErr(errOnDeleting)
	fmt.Println("the error on delete file",errOnDeleting)

}

func readFile(fileName string){
	dataByte ,err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println(string(dataByte))
	fmt.Printf("the type is %T\n",dataByte)
}

func checkNilErr(err error){
	if(err != nil){
		panic(err)
	}
}