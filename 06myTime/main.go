package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time is here")

	presentTime := time.Now()

	fmt.Println("THis is current time", presentTime)


	fmt.Println("THis is current time in format", presentTime.Format("2006-01-02"))
}