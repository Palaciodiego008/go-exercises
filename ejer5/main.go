package main

import "fmt"

var calculation func(int, int) int = func(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", calculation(5, 7))

	/// restar

	calculation = func(num1, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d \n", calculation(6, 4))

	calculation = func(num1, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Dividimos 5 / 5  = %d \n", calculation(5, 5))
	Operations()

	tableOf2 := 2
	MyTable := Table(tableOf2)

	for i := 1; i < 11; i++ {
		fmt.Println(MyTable())
	}

}

func Operations() {
	result := func() int {
		var a int = 23
		var b int = 13
		return a + b

	}
	fmt.Println(result())
}

func Table(value int) func() int {
	number := value
	sequence := 0
	return func() int {
		sequence += 1
		return number * sequence
	}
}
