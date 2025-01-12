package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array module")

	var playerList [3]string
	playerList[0] = "Hasan"

	fmt.Println("The playerList is", playerList[2])
	fmt.Println("The playerList length is", len(playerList))

	var vegList = [4]string{"Banana", "potato", "orange" }
	vegList[3] = "Custom orange"
	fmt.Println("VegList is", vegList)

	var playerListNum [3]int
	playerListNum[0] =2

	fmt.Println("The playerListNum is", playerListNum[2])
}