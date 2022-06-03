package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// This is making a HTTP GET request.
	// func http.Get(url string) (resp *http.Response, err error)
	// It returns a pointer to a http.Response and an error
	resp, err := http.Get("http://www.google.com")
	// Basic error handling.
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
