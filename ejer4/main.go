package main

import (
	"fmt"
)

func main() {
	// fmt.Println(one(5))
	// numero, estado := two(10)
	// fmt.Println(numero, estado)

	fmt.Println(Caculation(4, 54))
	fmt.Println(Caculation(2, 543))
	fmt.Println(Caculation(4, 24))
	fmt.Println(Caculation(1, 43, 5, 3, 54))

}

func one(number int) int {
	return number * 2

}

func two(number int) (int, bool) {
	if number == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func Caculation(number ...int) int {
	total := 0
	for item, num := range number {
		total = total + num
		fmt.Printf("\n item %d \n", item)
	}
	return total
}
