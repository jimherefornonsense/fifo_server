package main

import (
	"fmt"
	"os"
	"sync"
	"syscall"
)

var writePipes []string = []string{"/tmp/alltoP1", "/tmp/alltoP2"}
var readPipes []string = []string{"/tmp/allfromP1", "/tmp/allfromP2"}
var saying []string = []string{"Hello one!\n", "Hello two!\n"}
var clientNumbers int = 2

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func writeToPipe(wg *sync.WaitGroup, path string, msg string) {
	err := os.WriteFile(path, []byte(msg), 0)
	errorCheck(err)
	wg.Done()
}

func readFromPipe(path string) string {
	//fmt.Println(path)
	data, err := os.ReadFile(path)
	errorCheck(err)
	return string(data)
}

func main() {
	// Init pipes
	for i := 0; i < clientNumbers; i++ {
		syscall.Mkfifo(writePipes[i], 0644)
		syscall.Mkfifo(readPipes[i], 0666)
	}

	// Control logic
	var wg sync.WaitGroup
	var counter int = 0
	for counter != clientNumbers {
		for i := counter; i < clientNumbers; i++ {
			wg.Add(1)
			go writeToPipe(&wg, writePipes[i], saying[counter%clientNumbers])
		}

		// Read from pipe
		msg := readFromPipe(readPipes[counter%clientNumbers])
		fmt.Println(msg)
		wg.Wait()
		counter++
	}

	// Remove pipes
	for i := 0; i < clientNumbers; i++ {
		os.Remove(writePipes[i])
		os.Remove(readPipes[i])
	}
}
