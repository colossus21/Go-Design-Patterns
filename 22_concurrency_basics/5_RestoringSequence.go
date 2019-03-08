package main

import (
	"fmt"
	"strconv"
	"time"
)

// Reintroduce Lock Steps, this time, we are going to make a Fanned channel maintain its sequence

// Message struct

type Message struct {
	str string
	wait chan bool
}

func main() {
	Bob := MsgGenerator("Bob")
	Mike := MsgGenerator("Mike")
	// Fan In functions count in lockstep, ie. if Bob is locked, Ann still sends data
	FanIn := FanIn(Bob, Mike)

	// Run 10 iterations and print the received channel data
	for i:=0; i<10; i++ {
		msg1 := <-FanIn
		msg2 := <-FanIn
		fmt.Println(msg1.str)
		fmt.Println(msg2.str)
		<-msg1.wait // throwing away this value / stop blocking
		<-msg2.wait // throwing away this value / stop blocking
	}

}

// Also called Fan In function
func FanIn(c1, c2 <-chan Message) <-chan Message {
	cFanned := make(chan Message)
	go func() { for{ cFanned <- <-c1 } }()
	go func() { for{ cFanned <- <-c2 } }()
	return cFanned
}

// Message Generator

func MsgGenerator(name string) <-chan Message {
	c := make(chan Message)
	cWait := make(chan bool) //default: false
	go func() {
		for i:=0;;i++ {
			msg := Message{str:name+" "+strconv.Itoa(i), wait:cWait}
			c <- msg
			time.Sleep(time.Second)
			cWait <- true // false, true
		}
	}()
	return c
}