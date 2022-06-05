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
	// We want the body (html) of the response.

	// resp.Body is a type of io.ReadCloser.
	//
	// type ReadCloser interface {
	// 	  Reader
	// 	  Closer
	// }
	//
	// io.ReadCloser implements 2 interfaces; a type of io.Reader and io.Closer
	// io.Reader implements a function Read([]byte) (int, error)
	// io.Closer implements a function Close() (error)

	// As resp.Body implements io.ReadCloser interface, we know it must have
	// a function Read([]byte) (int, error)
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
