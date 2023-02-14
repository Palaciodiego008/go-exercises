package main

import (
	"fmt"
	"os"
)

func main() {
	file := "test.txt"
	// readFile
	f, err := os.Open(file)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	defer f.Close()

}

func examplePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f %w", r)
		}
	}()
	panic("panic")
}
