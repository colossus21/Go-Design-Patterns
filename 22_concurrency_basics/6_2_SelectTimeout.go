package main

import (
	"fmt"
	"math/rand"
	"time"
)

// We can minimized our code by using SELECT for FanIn functions

func selectFanIn(input1, input2, input3 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			case s := <-input3:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	centralServer := Gen("Central API")
	hrServer := Gen("HR API")
	devServer := Gen("Developer's API")

	fannedChan := selectFanIn(centralServer, hrServer, devServer)
	//timedOut := time.After(time.Second)
	for  {
		select {
		case c:=<-fannedChan:
			fmt.Println(c)
		//case <-timedOut: // Global Timeout
		case <-time.After(time.Second): // time.After(T) returns a CHANNEL, timeout for each message
			fmt.Println("Too slow!") // If no channels are found after within time T, this is executed.
			return //exits for
		}
	}
}

func Gen(name string) <-chan string {
	c := make(chan string)
	go func() { for {
		c <- "Receiving from" + name
		randTime := rand.Intn(3)
		time.Sleep(time.Duration(randTime) * time.Second)
	}
	}()
	return c
}



