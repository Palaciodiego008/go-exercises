package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go myNameSlow("Diego Dapo")
	fmt.Println("I'm here")
	var wait string
	fmt.Scanln(&wait)

}

func myNameSlow(name string) {

	letters := strings.Split(name, fmt.Sprintf("%c", 32))
	letters = append(letters, fmt.Sprintf("%c", 32))

	for _, letter := range letters {

		time.Sleep(11 + time.Millisecond)
		fmt.Println(letter)
	}

}
