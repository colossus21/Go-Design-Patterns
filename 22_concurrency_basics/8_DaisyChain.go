package main

import "fmt"

func stream(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	// ch 1 -> ch 2 -> ch 3 -> ch 4 -> ... -> ch n  [ channel n receives the value at the end ]
	n := 100 // No of goroutines
	initial := make(chan int)
	left := initial
	right := initial

	for i:=0;i<n;i++ {
		right = make(chan int) // Create a new channel on the right side
		go stream(left, right) // Value of left = right + 1
		left = right // Restart the loop, right channel is now the left channel
	}

	go func(c chan int) { c<-1 }(right) // Value of right is 1
	fmt.Println(<-initial)
}