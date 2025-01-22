package main

import (
	"fmt"
	"log"
	route "mongoAPI/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	r := route.Route()

	fmt.Println("Hey Server is starting...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening Server is on 4000 port")
}
