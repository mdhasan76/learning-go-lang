package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"site"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to bit more json of GoLang")
	encodeJson()
}

func encodeJson() {
	courseList := []Course{
		{"Typescript", 20, "webdevsimplified", "234", []string{"ts"}}, // way 1
		{"Go Lang", 20, "hiteshweb", "234", nil},                      // way 1
		// {Name: "Javascript", Price: 30, Platform: "threeDotAcademy", Password: "1234", Tags: []string{"js", "javascript", "ts"}}, // way 2
	}
	finalJsonConvert, err := json.MarshalIndent(courseList, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJsonConvert)
}
