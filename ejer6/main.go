package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arrayintToString := []string{}

	for _, value := range array {
		arrayintToString = append(arrayintToString, fmt.Sprintln(value))
	}

	fmt.Println(arrayintToString)
}
