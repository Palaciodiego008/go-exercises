package main

import (
	"fmt"
	"time"
)

type User struct {
	Id     int
	Name   string
	Date   time.Time
	Status bool
}

func main() {
	user := new(User)
	fmt.Println("user: ", user)
}
