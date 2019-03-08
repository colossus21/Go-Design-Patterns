package main

import "fmt"

// Each case is a communication
// BLOCKS until one of the send/receive operations is READY
// If MULTIPLE operations are ready, pick ONE of them at RANDOM


func main() {
	centralServer := ApiGenerator("Central API")
	hrServer := ApiGenerator("HR API")
	devServer := ApiGenerator("Developer's API")

	// Identify number of calls made
	callCounter := 0

	for i:=0;i<10;i++ {
		// Select is the switch statement for channels
		select {
		case <-centralServer:
			fmt.Println(<-centralServer)
			callCounter++
		case <-hrServer:
			fmt.Println(<-hrServer)
			callCounter++
		case <-devServer:
			fmt.Println(<-devServer)
			callCounter++
		default:
			fmt.Println("Didn't receive anything!")
			callCounter++
		}
	}

	// Prints number of calls
	fmt.Println("Called",callCounter,"times!")
}

func ApiGenerator(name string) <-chan string {
	c := make(chan string)
	go func() { for{ c<-"Receiving from"+name } }()
	return c
}
