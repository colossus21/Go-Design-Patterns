package main

import (
	"fmt"
	"sync"
	"time"
)

// Race condition: Two or more resources have potential to access same memory location

// Mutex can Lock and Unlock

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	// You can check if race condition exists by "go run -race 11_Mutex.go" command

	c1 := Counter{}
	c2 := Counter{}

	// Unhandled race condition
	for i:=0; i<10; i++ {
		go func(i int) {
			c1.value++
		}(i)
	}

	// Handled race condition
	for i:=0; i<10; i++ {
		go func(i int) {
			c2.Lock() // Locks other goroutines from accessing c2
			c2.value++
			c2.Unlock() // Unlocks access to c2 by other goroutines
		}(i)
	}


	c2.Lock() // Locks c2 at some point
	defer c2.Unlock() // Unlock memory access in the end
	time.Sleep(1)
	fmt.Println(c1.value)
	fmt.Println(c2.value)
}