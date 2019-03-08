package main

import (
	"fmt"
)

func main() {
	quit := make(chan bool) // When this bool is true, channels stop receiving
	c := FindNumber(25, quit)
	for {
		select {
		case s:=<-c:
			fmt.Println(s)
		case <-quit:
			return
		}
	}
}

func FindNumber(n int, q chan bool) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; ; i++ {
			c <- i
			if i == n {
				q <- true // q = true when number is found
			}
		}
	}()

	return c
}