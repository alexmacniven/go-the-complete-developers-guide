package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

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

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// The Copy function of the io package provides a way of taking a type
	// implementing the Writer interface and a type implementing the Reader
	// interface and copying the output of the reader to the input of the
	// writer. In a sense we're creating a channel from the Reader output to
	// the writer input.

	// As Copy accepts any type implementing the Writer interface, we can
	// write our own types to easily work with the built-in functionality.
	// This is where the power of interfaces comes into play.
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Printf("Just printed %v bytes.\n", len(p))
	return len(p), nil
}
