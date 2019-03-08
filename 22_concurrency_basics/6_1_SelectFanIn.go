package main

import (
	"fmt"
	"time"
)

// We can minimized our code by using SELECT for FanIn functions

func SelectFanIn(input1, input2, input3 <-chan string) <-chan string {
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
			time.Sleep(time.Second)
		}
	}()

	return c
}

func main() {
	centralServer := Generator("Central API")
	hrServer := Generator("HR API")
	devServer := Generator("Developer's API")

	fannedChan := SelectFanIn(centralServer, hrServer, devServer)

	for i:=0;i<10;i++ {
		fmt.Println(<-fannedChan)
	}
}

func Generator(name string) <-chan string {
	c := make(chan string)
	go func() { for{ c<-"Receiving from"+name } }()
	return c
}



