package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Command-line arguments are available with os.Args as a
	// string slice. Arguments start at slice index 1.
	if len(os.Args) == 1 {
		fmt.Println("Error: not enough command line arguments")
		os.Exit(1)
	}
	fn := os.Args[1]

	// The os.Open function returns a pointer to a os.File type
	// and an Error. The os.File type implements the Reader
	// interface.
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// As we know, os.Stdout is of os.File type which implements
	// the Writer interface.
	//
	// We can use io.Copy to read from the file and write to Stdout.
	io.Copy(os.Stdout, f)
}
