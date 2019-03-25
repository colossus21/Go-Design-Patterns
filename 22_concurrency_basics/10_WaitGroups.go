package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // Wait for two Goroutines to finish

	// Any two of the following goroutines will be executed

	go func() {
		fmt.Println("Calling from Func 1")
		wg.Done() // Finishes the goroutine, subtracts 1 from wg counter
	}()

	go func() {
		fmt.Println("Calling from Func 2")
		wg.Done() // Finishes the goroutine, subtracts 1 from wg counter
	}()

	go func() {
		fmt.Println("Calling from Func 3")
		wg.Done() // Finishes the goroutine, subtracts 1 from wg counter
	}()

	go func() {
		fmt.Println("Calling from Func 4")
		wg.Done() // Finishes the goroutine, subtracts 1 from wg counter
	}()

	// wg.Done() is equivalent to wg.Add(-1)

	wg.Wait() // Wait for the necessary time for 2 goroutines to finish
}