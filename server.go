package main

import (
	"fmt"
	"os"
	"syscall"
)

var writePipes []string = []string{"/tmp/alltoP1", "/tmp/alltoP2"}
var readPipes []string = []string{"/tmp/allfromP1", "/tmp/allfromP2"}
var saying []string = []string{"Hello one!", "Hello two!"}
var clientNumbers int = 2

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Init pipes
	for i := 0; i < clientNumbers; i++ {
		syscall.Mkfifo(writePipes[i], 0644)
		syscall.Mkfifo(readPipes[i], 0666)
	}

	// Control logic
	var counter int = 0
	for counter != clientNumbers {
		// Write to pipes
		for i := 0; i < clientNumbers; i++ {
			err := os.WriteFile(writePipes[i], []byte(saying[counter%clientNumbers]), 0)
			errorCheck(err)
		}

		// Read from pipe
		data, err := os.ReadFile(readPipes[counter%clientNumbers])
		errorCheck(err)
		fmt.Println(string(data))
		counter++
	}

	// Remove pipes
	for i := 0; i < clientNumbers; i++ {
		os.Remove(writePipes[i])
		os.Remove(readPipes[i])
	}
}
