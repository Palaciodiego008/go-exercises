package main

import (
	"fmt"
)

func main() {
	var i int = 0
RUTINA:
	for i < 8 {
		if i == 4 {
			i += i + 3
			fmt.Println("voy a RUTINA sumando a 2 i ")
			goto RUTINA
		}
		fmt.Printf("valor de i: %d\n", i)
		i++
	}

}
