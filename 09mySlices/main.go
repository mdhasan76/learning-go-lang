package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("Welcome to Slice of GO lang")
	//  num := 19
	// // var nameList = []string{}
	// fmt.Printf("The name list is %T \n", num)

	// _______________Array
	var bookNameArr = [3]string{}
	fmt.Println("The book list", bookNameArr)
	bookNameArr[0] = "Ar Rahikum Makhtum"
	bookNameArr[1] = "Jannat Jahannam"
	// bookNameArr[2] = "Powerful focus"
	fmt.Println("list", bookNameArr)
	fmt.Println("Length", len(bookNameArr))
	fmt.Println("Capacity", cap(bookNameArr))
	
	// _______________Slice
	var nameList = []string{"Hasan", "robin"}
	fmt.Println("The nameList is", nameList)
	nameList = append(nameList, "Yasin")
	fmt.Println("The nameList is", nameList)
	fmt.Printf("The nameList is %T \n", nameList)
	fmt.Println(len(nameList))
	nameList = append(nameList[:1])
	fmt.Println(len(nameList))
	fmt.Println(nameList)


	// create slice 
	slice2 := make([]int, 4)
	fmt.Println(slice2)
	fmt.Println(sort.IntsAreSorted(slice2))
	fmt.Println("length", len(slice2))
	slice2[0] = 90
	slice2[1] = 60
	fmt.Println(sort.IntsAreSorted(slice2))
	slice2[2] = 45
	slice2[3] = 102
	fmt.Println(slice2)
	 sort.Ints(slice2)
	 fmt.Println(slice2)
	 fmt.Println(sort.IntsAreSorted(slice2))
	fmt.Println(slice2)
}