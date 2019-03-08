package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Goroutine execution depends on the main function's scope
	// Define channel				:		c := make(chan int)
	// Setting value of a channel	:		c <- 10
	// Getting value from channel 	:		val := <-c

	// Create channel
	c := make(chan string)

	// Buffered channel removes synchronization
	cb := make(chan string, 4)


	// Running goroutine that streams increasing values of i
	go generator(c)
	go generatorb(cb)

	// Run 10 iterations and print the received channel data
	for i:=0; i<10; i++ {
		ch1 := <-c
		ch2 := <-cb
		// If ch1 is locked, ch2 doesn't continue
		fmt.Println("[channel#1]", ch1)
		fmt.Println("[channel#2]", ch2)
	}
}

func generator(c chan string) {
	// Forever loop generates increasing values of "i"; sends the value to the channel
	for i:=0;;i++ {
		c <- "N:"+strconv.Itoa(i)
	}
}

func generatorb(c chan string) {
	// Forever loop generates increasing values of "i"; sends the value to the channel
	for i:=0;;i++ {
		c <- "N:"+strconv.Itoa(i)
	}
}