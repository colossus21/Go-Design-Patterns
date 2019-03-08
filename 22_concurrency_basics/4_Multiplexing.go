package main

import (
	"fmt"
)

func main() {
	Bob := Say("Bob")
	Mike := Say("Mike")
	// Fan In functions count in lockstep, ie. if Bob is locked, Ann still sends data
	FanIn := Multiplexer(Bob, Mike)

	// Run 10 iterations and print the received channel data
	for i:=0; i<10; i++ {
		fmt.Println(<-FanIn)
	}


	// Let's restore the sequence

}

// Also called Fan In function
func Multiplexer(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for{ c <- <-c1 } }()
	go func() { for{ c <- <-c2 } }()
	return c
}

func Say(name string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- name + " says hi!"
		}
	}()
	return c
}