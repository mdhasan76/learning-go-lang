package main

import "fmt"

func main() {
	fmt.Println("Welcome in go map tutorial")
	laptops := make(map[string]string)
	laptops["hp"] = "HP elite"
	laptops["mac"] = "MackbookPro"

	delete(laptops, "hp")
	fmt.Println(laptops)

	var ram = make(map[int]string)
	ram[1] = "Winston"
	ram[2] = "Corsiar"
	fmt.Println(ram)

	forLoopTest(10)
}

func forLoopTest(rang int){
	for i := 1; i < rang; i++{
		fmt.Println(i)
	}
}