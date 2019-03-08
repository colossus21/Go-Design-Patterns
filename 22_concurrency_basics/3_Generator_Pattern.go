package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumberGenerator() <-chan int {
	// Create channel
	c := make(chan int)
	// Dynamic seed
	rand.Seed(time.Now().UnixNano())
	// Start goroutine to update channel
	go func() {
		for {
			c <- rand.Intn(100)
		}
	}()

	// Return channel
	return c
}

func main() {
	// Receive channel
	t := RandomNumberGenerator()
	// Try to receive the values
	for i:=0; i<10; i++ {
		fmt.Println("Number:", <-t)
	}
}
