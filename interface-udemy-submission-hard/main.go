package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Get argument for file name
	fileName := os.Args[1]

	// Open file name by current argument
	fileData, err := os.Open(fileName)

	// If there is error
	if err != nil {
		log.Fatal(err)
	} else {
		// If there is no error
		io.Copy(os.Stdout, fileData)
	}
}
