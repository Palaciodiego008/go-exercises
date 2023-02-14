package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne := make(chan time.Duration)

	go bucle(channelOne)
	fmt.Println("Llego hasta aqui")

	msg := <-channelOne
	fmt.Println(msg)
}

func bucle(channel chan time.Duration) {
	init := time.Now()

	for i := 0; i < 100000000; i++ {
	}

	final := time.Now()

	channel <- final.Sub(init)
}
