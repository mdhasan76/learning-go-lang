package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Maths in Go lang from hasan")
	rand.Seed(4 * time.Now().UnixNano())
	var randomNum int = rand.Intn(86)
	fmt.Println(randomNum)
}
