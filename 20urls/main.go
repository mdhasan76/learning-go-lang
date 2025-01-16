package main

import (
	"fmt"
	"net/url"
)

const link = "https://www.facebook.com:3000/reel?courseName=golang&page=3"

func main() {
	fmt.Println("Welcome to go urls")
	formate, err := url.Parse(link)
	if err != nil {
		panic(err)
	}
	fmt.Println(formate.Host)
	// fmt.Println(formate.Path)
	// fmt.Println(formate.Scheme)
	// fmt.Println(formate.User)
	// fmt.Println(formate.RawQuery)
	// fmt.Println(formate.RawFragment)
	// fmt.Println(formate.Port())

	qParams := formate.Query()
	fmt.Println(qParams)
	fmt.Println("the value is", qParams["courseName"])
	fmt.Printf("the type of qParams is %T\n", qParams)
	for _, val := range qParams {
		fmt.Println("param is:", val)
	}
	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "www.facebook.com",
		Path:     "/reel",
		RawQuery: "userName=Hasan",
		RawPath:  "hello",
	}
	fmt.Println(partsOfURL.String())
}
