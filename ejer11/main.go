package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	readFile()
	createCSV()
	createFile()
}

func readFile() {
	file, err := ioutil.ReadFile("./some.log")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

}

func createFile() {
	file, err := os.Create("./newSome.log")
	file.WriteString("Hello World pal")

	if err != nil {
		fmt.Println(err)
	}

	file.Close()

}

func createCSV() {

	file, err := os.Create("./newSome.csv")
	file.WriteString("Hello World pal")

	if err != nil {
		fmt.Println(err)
	}

	file.Close()
}
