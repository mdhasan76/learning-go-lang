package main

import (
	"fmt"
	"net/http"
	"sync"
)

var webLinkList = []string{"hello"}
var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

func main() {
	fmt.Println("Welcome to go routines")
	websiteList := []string{
		// "https://lco.dev",
		"https://github.com",
		"https://github.com/mdhasan76",
		"https://google.com",
		"https://fb.com",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(websiteList)
	// go greeting("Subahan-allah")
	// greeting("Alhadolillah")
}

// func greeting(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(5 * time.Microsecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	// time.Sleep(5 * time.Second)
	if err != nil {
		fmt.Println("OOPS error occurs in endpoint")
	}
	mut.Lock()
	webLinkList = append(webLinkList, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for  %s\n", res.StatusCode, endpoint)
}
