package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
		fileContents, err := ioutil.ReadFile("first-post.txt")
		if err != nil {
				// A common use of `panic` is to abort if a function returns an error
				// value that we don’t know how to (or want to) handle. This example
				// panics if we get an unexpected error when creating a new file.
				panic(err)
		}
		fmt.Print(string(fileContents))

		bytesToWrite := []byte("hello\ngo\n")
        err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
        if err != nil {
            panic(err)
        }
}
