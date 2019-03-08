package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan string) // When this bool is true, channels stop receiving
	c := findNumber(quit)
	for i:=0; i<10; i++{
		v := <-c
		fmt.Println(v)
	}
	quit <- "Done!" // Send a signal to quit
	fmt.Println("Final message:",<-quit)
}

func findNumber(q chan string) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; ; i++ {
			select {
			case c<-i:
				time.Sleep(time.Second/4)
			case <-q: // Receives quit signal
				fmt.Println("Doing stuff like backups, storing log!!") // Perform necessary operations!
				q <- "Safely Executed!!" // Send final message!
				return
			}
		}
	}()

	return c
}